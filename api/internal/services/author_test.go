package services

import (
	"api/api"
	"api/internal/db"
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

var authorSvc AuthorService
var mockAuthorQueries *MockAuthorQueries
var mockValidator *MockValidator
var sqlMock sqlmock.Sqlmock

func setupTest(t testing.TB) func(t testing.TB) {
	log.Println("setup test")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorQueries = NewMockAuthorQueries(ctrl)
	mockValidator = NewMockValidator(ctrl)

	var database *sql.DB
	var err error
	database, sqlMock, err = sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%authorSvc' was not expected when opening a stub database connection", err)
	}

	authorSvc = NewAuthorService(database, mockAuthorQueries, mockValidator, &log.Logger{})

	return func(t testing.TB) {
		log.Println("teardown test")
		database.Close()
	}
}

func Test_authorService_Create(t *testing.T) {
	ctx := context.TODO()
	author := api.AuthorBase{
		LastName:   "Last",
		FirstName:  "First",
		MiddleName: "Middle",
	}
	expected := int64(1)

	teardownTest := setupTest(t)
	defer teardownTest(t)
	mockValidator.EXPECT().StructCtx(ctx, author).Return(nil).Times(1)
	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()
	mockAuthorQueries.EXPECT().WithTx(gomock.Any()).Return(mockAuthorQueries)
	mockAuthorQueries.EXPECT().GetUniqueAuthor(ctx, gomock.Any()).Return(db.Author{}, nil)
	mockAuthorQueries.EXPECT().CreateAuthor(ctx, gomock.Any()).Return(expected, nil)

	s := authorSvc
	got, err := s.Create(ctx, author)
	if err != nil {
		t.Errorf("Create() error = %v", err)
		return
	}
	if got != expected {
		t.Errorf("Create() got = %v, want %v", got, expected)
	}
}
