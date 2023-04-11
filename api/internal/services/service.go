package services

import (
	"api/internal/db"
	"api/internal/model"
	"context"
	"database/sql"
	"fmt"
	"sync"
)

// AuthorService interface represents all the methods that a service needs to manage authors.
type AuthorService interface {
	// Create a new author.
	Create(ctx context.Context, author model.AuthorBase) (int64, error)
}

// Logger interface represents the method required for a logger.
type Logger interface {
	Printf(format string, args ...interface{})
}

// Validator interface represents the methods required for a validator.
type Validator interface {
	Struct(interface{}) error
	StructCtx(ctx context.Context, s interface{}) (err error)
}

type authorService struct {
	db        *sql.DB
	queries   *db.Queries
	mutex     *sync.Mutex
	logger    Logger
	validator Validator
}

func NewAuthorService(db *sql.DB, queries *db.Queries, validator Validator, logger Logger) AuthorService {
	return &authorService{
		db:        db,
		queries:   queries,
		mutex:     &sync.Mutex{},
		validator: validator,
		logger:    logger,
	}
}

func (s *authorService) Create(ctx context.Context, author model.AuthorBase) (int64, error) {

	if err := s.validator.StructCtx(ctx, author); err != nil {
		return -1, NewValidationErr("Author", err.Error())
	}

	var id int64
	err := s.transaction(func(q *db.Queries) error {
		authorDb, err := q.GetUniqueAuthor(ctx, toGetUniqueAuthorParams(author))
		if err != nil && err != sql.ErrNoRows {
			s.logger.Printf("Create - query call error: %s", err)
			return fmt.Errorf("create new author error: %w", err)
		}
		if authorDb != (db.Author{}) {
			return NewDuplicateErr("Author", fmt.Sprintf("%v %v %v", author.FirstName, author.MiddleName, author.LastName))
		}

		id, err = q.CreateAuthor(ctx, toGetCreateAuthorParams(author))
		return nil
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *authorService) transaction(action func(queries *db.Queries) error) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("error statring a transaction: %w", err)
	}
	defer tx.Rollback()
	q := s.queries.WithTx(tx)
	err = action(q)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error commiting the transaction : %w", err)
	}
	return nil
}

func toGetCreateAuthorParams(author model.AuthorBase) db.CreateAuthorParams {
	params := db.CreateAuthorParams{
		LastName:   author.LastName,
		FirstName:  author.FirstName,
		MiddleName: author.MiddleName,
	}
	return params
}

func toGetUniqueAuthorParams(author model.AuthorBase) db.GetUniqueAuthorParams {
	params := db.GetUniqueAuthorParams{
		LastName:   author.LastName,
		FirstName:  author.FirstName,
		MiddleName: author.MiddleName,
	}
	return params
}
