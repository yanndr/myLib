-- name: CreateAuthor :one
INSERT INTO authors (last_name, first_name, middle_name)
VALUES (?, ?, ?)
RETURNING id;

-- name: GetUniqueAuthor :one
SELECT * FROM authors WHERE last_name=? AND first_name=? AND middle_name=?;

-- name: GetAuthorById :one
SELECT * FROM authors WHERE id = ?;

-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = ?;

-- name: GetAllAuthors :many
SELECT * FROM authors ORDER BY last_name;

-- name: GetAllAuthorsWithName :many
SELECT * FROM authors WHERE last_name =? ORDER BY last_name;

-- name: UpdateAuthor :exec
UPDATE authors
SET last_name=?,
    first_name=?,
    middle_name=?
WHERE id = ?;