// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: authors.sql

package db

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (last_name, first_name, middle_name)
VALUES (?, ?, ?)
RETURNING id
`

type CreateAuthorParams struct {
	LastName   string
	FirstName  string
	MiddleName string
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.LastName, arg.FirstName, arg.MiddleName)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = ?
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAllAuthors = `-- name: GetAllAuthors :many
SELECT id, first_name, last_name, middle_name FROM authors ORDER BY last_name
`

func (q *Queries) GetAllAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, getAllAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.MiddleName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllAuthorsWithName = `-- name: GetAllAuthorsWithName :many
SELECT id, first_name, last_name, middle_name FROM authors WHERE last_name =? ORDER BY last_name
`

func (q *Queries) GetAllAuthorsWithName(ctx context.Context, lastName string) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, getAllAuthorsWithName, lastName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.MiddleName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAuthorById = `-- name: GetAuthorById :one
SELECT id, first_name, last_name, middle_name FROM authors WHERE id = ?
`

func (q *Queries) GetAuthorById(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthorById, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
	)
	return i, err
}

const getUniqueAuthor = `-- name: GetUniqueAuthor :one
SELECT id, first_name, last_name, middle_name FROM authors WHERE last_name=? AND first_name=? AND middle_name=?
`

type GetUniqueAuthorParams struct {
	LastName   string
	FirstName  string
	MiddleName string
}

func (q *Queries) GetUniqueAuthor(ctx context.Context, arg GetUniqueAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, getUniqueAuthor, arg.LastName, arg.FirstName, arg.MiddleName)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
	)
	return i, err
}
