package api

import (
	"crypto/md5"
	"encoding/gob"
	"fmt"
)

const (
	AuthorsPath = "/authors"
	BooksPath   = "/books"
)

// Information represents the information about the API.
type Information struct {
	APIVersion string `json:"api_version"`
}

type Serializable interface {
	Serialize() (string, error)
}

func Serialize(data interface{}) (string, error) {
	writer := md5.New()
	e := gob.NewEncoder(writer)
	err := e.Encode(data)
	if err != nil {
		return "", fmt.Errorf(`failed gob Encode :%w`, err)
	}
	return fmt.Sprintf("%x", writer.Sum(nil)), nil
}

// AuthorBase represents the base information for an author.
type AuthorBase struct {
	LastName   string `json:"last_name" validate:"required"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
}

// Author represent an author.
type Author struct {
	ID int64 `json:"ID"`
	AuthorBase
}

// CreateUpdateAuthorRequest represents the structure used for creating or updating an Author
type CreateUpdateAuthorRequest struct {
	AuthorBase
}

// PatchAuthorRequest represents the structure used for creating or partially updating an Author
type PatchAuthorRequest struct {
	AuthorBase
	ModifiedLastName   bool `json:"modified_last_name"`
	ModifiedFirstName  bool `json:"modified_first_name"`
	ModifiedMiddleName bool `json:"modified_middle_name"`
}

func (a Author) Serialize() (string, error) {
	return Serialize(a)
}
