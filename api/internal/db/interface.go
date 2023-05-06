package db

import (
	"context"
	"database/sql"
)

type dbQueries interface {
	WithTx(tx *sql.Tx) *Queries
}

type AuthorQueries interface {
	dbQueries
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (int64, error)
	DeleteAuthor(ctx context.Context, id int64) error
	GetAllAuthors(ctx context.Context, arg GetAllAuthorsParams) ([]Author, error)
	GetAllAuthorsWithName(ctx context.Context, lastName string) ([]Author, error)
	GetAuthorById(ctx context.Context, id int64) (Author, error)
	GetUniqueAuthor(ctx context.Context, arg GetUniqueAuthorParams) (Author, error)
	UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) error
}

type AllQueries interface {
	AuthorQueries
}
