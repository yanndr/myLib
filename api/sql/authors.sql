-- name: CreateAuthor :one
INSERT INTO authors (last_name, first_name, middle_name)
VALUES (?, ?, ?)
RETURNING id;

-- name: GetUniqueAuthor :one
SELECT * FROM authors WHERE last_name=? AND first_name=? AND middle_name=?;

-- name: GetAuthorById :one
SELECT * FROM authors WHERE id = ?;