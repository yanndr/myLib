package main

import (
	"api/internal/db"
	"api/internal/endpoints"
	"api/internal/middlewares"
	"api/internal/services"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	_ "github.com/mattn/go-sqlite3"
)

var Version = "0.1-dev"

const (
	dbName  = "books.db"
	appName = "myLib"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "port default 8080")
	flag.Parse()

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	appConfigDir := path.Join(home, "."+appName)

	if err := run(port, appConfigDir, dbName); err != nil {
		log.Fatal(err)
	}
}

func run(port int, configDir, dbName string) error {
	if _, err := os.Stat(configDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(configDir, 0700)
		if err != nil {
			return err
		}
	}

	database, err := db.OpenDatabase(path.Join(configDir, dbName))
	if err != nil {
		return err
	}
	queries := db.New(database)
	authSvcLogger := log.New(os.Stderr, "API - AuthorService -", log.LstdFlags)
	authSvc := services.NewAuthorService(database, queries, validator.New(), authSvcLogger)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Method(http.MethodGet, "/metrics", promhttp.Handler())
	r.Get("/", endpoints.RootResponse)
	createRoutes(r, endpoints.NewV1Route(Version, authSvc))

	return startServer(port, r, err)
}

const socketPath = "/tmp/mylib.sock"

func startServer(port int, r http.Handler, err error) error {
	server := &http.Server{Addr: fmt.Sprintf(":%v", port), Handler: r}

	socket, err := net.Listen("unix", socketPath)
	if err != nil {
		return err
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	go func() {
		<-sig
		log.Println("Stopping the server...")
		os.Remove(socketPath)
		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	go func() {
		fmt.Printf("Server started, listening on socket: %v\n", socketPath)
		if err := server.Serve(socket); err != nil {
			fmt.Println("cannot listen on socket")
		}
	}()

	fmt.Printf("Server started, listening on port: %v\n", port)
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
	return nil
}

func createRoutes(router chi.Router, endpoint *endpoints.Route) {
	router.Route(endpoint.Pattern, func(r chi.Router) {
		for method, action := range endpoint.Actions {
			r.With(middlewares.PrometheusMiddleware).Method(method, "/", http.HandlerFunc(endpoints.Handle(action)))
		}
		if endpoint.SubRoutes != nil {
			for _, e := range endpoint.SubRoutes {
				createRoutes(r, e)
			}
		}
	})
}
