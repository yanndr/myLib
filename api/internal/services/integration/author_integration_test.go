package integration

import (
	"api/api"
	"api/internal/db"
	"api/internal/services"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"reflect"
	"testing"
)

var authorSvc services.AuthorService

func setup(t testing.TB) func(t testing.TB) {
	log.Println("setup test")

	database, err := db.OpenInMemoryDatabase()
	if err != nil {
		t.Fatalf("an error %v was not expected when opening a stub database connection", err)
	}
	queries := db.New(database)
	authorSvc = services.NewAuthorService(database, queries, validator.New(), &log.Logger{})

	return func(t testing.TB) {
		log.Println("teardown test")
		database.Close()
	}
}

func Test_authorService_Create(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	ctx := context.TODO()
	author := api.AuthorBase{
		LastName:   "Last",
		FirstName:  "First",
		MiddleName: "Middle",
	}

	id, err := authorSvc.Create(ctx, author)
	if err != nil {
		t.Errorf("Create() error = %v", err)
		return
	}

	if id != 1 {
		t.Errorf("expected id=1, got id=%v", id)
	}

	a, err := authorSvc.GetById(ctx, 1)
	if err != nil {
		t.Errorf("Create() error = %v", err)
		return
	}
	if !reflect.DeepEqual(a.AuthorBase, author) {
		t.Errorf("Create() got = %v, want %v", a.AuthorBase, author)
	}
}

func Benchmark_authorService_Create(b *testing.B) {
	teardown := setup(b)
	defer teardown(b)

	for i := 0; i < b.N; i++ {

		ctx := context.TODO()
		author := api.AuthorBase{
			LastName:   fmt.Sprintf("last%v", i),
			FirstName:  "First",
			MiddleName: "Middle",
		}

		_, err := authorSvc.Create(ctx, author)
		if err != nil {
			b.Errorf("Create() error = %v", err)
			return
		}
	}
}
