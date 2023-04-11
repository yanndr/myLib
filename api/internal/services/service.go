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
	StructCtx(ctx context.Context, s interface{}) (err error)
}

type service struct {
	db        *sql.DB
	queries   *db.Queries
	mutex     *sync.Mutex
	logger    Logger
	validator Validator
}

func newService(db *sql.DB, queries *db.Queries, validator Validator, logger Logger) service {
	return service{
		db:        db,
		queries:   queries,
		mutex:     &sync.Mutex{},
		logger:    logger,
		validator: validator,
	}
}

func (s *service) transaction(action func(queries *db.Queries) error) error {
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
