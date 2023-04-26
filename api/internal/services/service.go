package services

import (
	"api/internal/db"
	"context"
	"database/sql"
	"fmt"
	"sync"
)

// Logger interface represents the method required for a logger.
type Logger interface {
	Printf(format string, args ...interface{})
}

// Validator interface represents the methods required for a validator.
type Validator interface {
	Struct(interface{}) error
	StructCtx(context.Context, interface{}) error
}

type service struct {
	db        *sql.DB
	queries   db.AuthorQueries
	mutex     *sync.Mutex
	logger    Logger
	validator Validator
}

func newService(db *sql.DB, queries db.AuthorQueries, validator Validator, logger Logger) service {
	return service{
		db:        db,
		queries:   queries,
		mutex:     &sync.Mutex{},
		logger:    logger,
		validator: validator,
	}
}

func (s *service) transaction(action func(queries db.AllQueries) error) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("error statring a transaction: %w", err)
	}
	q := s.queries.WithTx(tx)
	err = action(q)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("update drivers: unable to rollback: %v with error: %w", rollbackErr, err)
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error commiting the transaction : %w", err)
	}
	return nil
}
