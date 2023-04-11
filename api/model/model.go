package model

// APIInformation represents the information about the API.
type APIInformation struct {
	APIVersion string `json:"api_version"`
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
