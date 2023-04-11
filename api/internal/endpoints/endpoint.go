package endpoints

import (
	"api/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	AuthorsPath = "/authors"
	BooksPath   = "/books"
	apiV1       = "/v1"
)

var NotImplementedErr = fmt.Errorf("not implemented")

// EndpointHandler represents a request method that returns a model.APIResponse.
type EndpointHandler func(r *http.Request) (model.APIResponse, error)

// Route represents a specific route of a server. It contains all the EndpointHandler mapped
// to their http.methods and all the sub routes
type Route struct {
	Pattern   string
	Actions   map[string]EndpointHandler
	SubRoutes []*Route
}

// NewV1Route creates the V1 route of the API
func NewV1Route() *Route {
	return &Route{
		Pattern: apiV1,
		Actions: map[string]EndpointHandler{
			http.MethodGet: notImplementedHandler,
		},
		SubRoutes: []*Route{
			&authorsEndpoint,
		},
	}
}

// Handle responds to an HTTP request,  executes the action EndpointHandler
// and writes the result to the http.ResponseWriter
func Handle(action EndpointHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response, err := action(r)
		if err != nil {
			if errors.Is(err, NotImplementedErr) {
				w.WriteHeader(http.StatusNotImplemented)
			}
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(response.StatusCode)
		_ = json.NewEncoder(w).Encode(response)
	}
}

func notImplementedHandler(_ *http.Request) (model.APIResponse, error) {
	return model.APIResponse{}, NotImplementedErr
}

var authorsEndpoint = Route{
	Pattern: AuthorsPath,
	Actions: map[string]EndpointHandler{
		http.MethodGet:  notImplementedHandler,
		http.MethodPost: notImplementedHandler,
	},
	SubRoutes: []*Route{&authorEndpoint},
}

var authorEndpoint = Route{
	Pattern: "/{id}",
	Actions: map[string]EndpointHandler{
		http.MethodGet:    notImplementedHandler,
		http.MethodPut:    notImplementedHandler,
		http.MethodPatch:  notImplementedHandler,
		http.MethodDelete: notImplementedHandler,
	},
	SubRoutes: []*Route{&authorBooksEndpoint},
}

var authorBooksEndpoint = Route{
	Pattern: BooksPath,
	Actions: map[string]EndpointHandler{
		http.MethodGet: notImplementedHandler,
	},
}
