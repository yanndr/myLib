package services

import (
	"api/api"
	"api/internal/db"
	"context"
	"database/sql"
	"fmt"
)

// AuthorService interface represents all the methods that a service needs to manage authors.
type AuthorService interface {
	// Create a new author.
	Create(ctx context.Context, author api.AuthorBase) (int64, error)
	// GetById returns a model.Author based on its id.
	GetById(ctx context.Context, id int64) (api.Author, error)
	// Delete deletes an author.
	Delete(ctx context.Context, id int64) error
	// GetAll returns the list of all authors.
	GetAll(ctx context.Context, nameFilter string) ([]api.Author, error)
	// Update updates an author.
	Update(ctx context.Context, id int64, author api.AuthorBase) error
	// PartialUpdate partially updates an author.
	PartialUpdate(ctx context.Context, id int64, patchRequest api.PatchAuthorRequest) error
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
	service
}

func NewAuthorService(db *sql.DB, queries *db.Queries, validator Validator, logger Logger) AuthorService {
	return &authorService{
		service: newService(db, queries, validator, logger),
	}
}

func (s *authorService) Create(ctx context.Context, author api.AuthorBase) (int64, error) {

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

func (s *authorService) GetById(ctx context.Context, id int64) (api.Author, error) {
	var author api.Author
	err := s.transaction(func(q *db.Queries) error {

		dbAuthor, err := q.GetAuthorById(ctx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return NewNotFoundErr("Author", id)
			}
			s.logger.Printf("GetById - query call error %s", err)
			return fmt.Errorf("get author by id error:  %w", err)
		}
		author = toApiAuthor(dbAuthor)

		return nil
	})
	if err != nil {
		return api.Author{}, err
	}

	return author, nil
}

func (s *authorService) Delete(ctx context.Context, id int64) error {
	err := s.transaction(func(q *db.Queries) error {
		_, err := q.GetAuthorById(ctx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return NewNotFoundErr("Author", id)
			}

			s.logger.Printf("Delete - query call error, %s", err)
			return fmt.Errorf("get author by id error: %w", err)
		}

		err = q.DeleteAuthor(ctx, id)
		if err != nil {
			s.logger.Printf("Delete - query call error, %s", err)
			return fmt.Errorf("deleting author %v error: %w", id, err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *authorService) GetAll(ctx context.Context, nameFilter string) ([]api.Author, error) {
	var authors []db.Author
	var err error
	if nameFilter != "" {
		authors, err = s.queries.GetAllAuthorsWithName(ctx, nameFilter)
	} else {
		authors, err = s.queries.GetAllAuthors(ctx)
	}

	if err != nil && err != sql.ErrNoRows {
		s.logger.Printf("GetAll - query call error, %s\n", err)
		return nil, fmt.Errorf("retrieving all authors error: %w", err)
	}
	result := make([]api.Author, len(authors))
	for i, a := range authors {
		result[i] = toApiAuthor(a)
	}
	return result, nil
}

func (s *authorService) Update(ctx context.Context, id int64, author api.AuthorBase) error {
	return s.update(ctx, id, author, func(dbAuthor *db.Author) {
		dbAuthor.LastName = author.LastName
		dbAuthor.FirstName = author.FirstName
		dbAuthor.MiddleName = author.MiddleName
	})
}

func (s *authorService) PartialUpdate(ctx context.Context, id int64, patchRequest api.PatchAuthorRequest) error {
	return s.update(ctx, id, patchRequest.AuthorBase, func(dbAuthor *db.Author) {
		if patchRequest.ModifiedLastName {
			dbAuthor.LastName = patchRequest.LastName
		}
		if patchRequest.ModifiedFirstName {
			dbAuthor.FirstName = patchRequest.FirstName
		}
		if patchRequest.ModifiedMiddleName {
			dbAuthor.MiddleName = patchRequest.MiddleName
		}
	})
}

func (s *authorService) update(ctx context.Context, id int64, author api.AuthorBase, updateFunc func(*db.Author)) error {
	err := s.transaction(func(q *db.Queries) error {
		dbAuthor, err := q.GetAuthorById(ctx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return NewNotFoundErr("Author", id)
			}
			s.logger.Printf("update -  query call error: %s", err)
			return fmt.Errorf("get author by id error: %w", err)
		}

		updateFunc(&dbAuthor)
		apiAuthor := toApiAuthor(dbAuthor)

		err = s.validator.StructCtx(ctx, apiAuthor)
		if err != nil {
			return NewValidationErr("Author", err.Error())
		}

		err = q.UpdateAuthor(ctx, toUpdateAuthorParams(apiAuthor))

		existing, err := q.GetUniqueAuthor(ctx, toGetUniqueAuthorParams(apiAuthor.AuthorBase))
		if err != nil {
			if err != sql.ErrNoRows {
				s.logger.Printf("update -  query call error: %s\n", err)
				return fmt.Errorf("update Author error: %w", err)
			}
		}
		if existing != (db.Author{}) && existing.ID != apiAuthor.ID {
			return NewDuplicateErr("Author", fmt.Sprintf("%v %v %v", author.FirstName, author.MiddleName, author.LastName))
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func toUpdateAuthorParams(author api.Author) db.UpdateAuthorParams {
	params := db.UpdateAuthorParams{
		LastName:   author.LastName,
		FirstName:  author.FirstName,
		MiddleName: author.MiddleName,
		ID:         author.ID,
	}
	return params
}

func toGetCreateAuthorParams(author api.AuthorBase) db.CreateAuthorParams {
	params := db.CreateAuthorParams{
		LastName:   author.LastName,
		FirstName:  author.FirstName,
		MiddleName: author.MiddleName,
	}
	return params
}

func toGetUniqueAuthorParams(author api.AuthorBase) db.GetUniqueAuthorParams {
	params := db.GetUniqueAuthorParams{
		LastName:   author.LastName,
		FirstName:  author.FirstName,
		MiddleName: author.MiddleName,
	}
	return params
}

func toApiAuthor(author db.Author) api.Author {
	return api.Author{
		ID: author.ID,
		AuthorBase: api.AuthorBase{
			LastName:   author.LastName,
			FirstName:  author.FirstName,
			MiddleName: author.MiddleName,
		},
	}
}
