package controllers

import (
	"api/internal/api"
	"api/internal/model"
	"api/internal/services"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// AuthorController represents all the endpoints.EndpointHandler related to Author.
type AuthorController struct {
	AuthorService services.AuthorService
	BasePath      string
}

// Create is the endpoint action for the HTTP POST method for creating a new author.
func (c *AuthorController) Create(r *http.Request) (model.APIResponse, error) {
	var author model.CreateUpdateAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		return model.APIResponse{}, api.NewBadFormatErr(err.Error())
	}
	if author == (model.CreateUpdateAuthorRequest{}) {
		return model.APIResponse{}, api.NewBadFormatErr("the author is empty")
	}

	id, err := c.AuthorService.Create(r.Context(), author.AuthorBase)
	if err != nil {
		return handleServiceError(err)
	}

	location := fmt.Sprintf("%v/%v", c.BasePath, id)
	return model.NewCreatedResponse(location), nil
}

func handleServiceError(err error) (model.APIResponse, error) {
	if errors.As(err, &services.DuplicateErr{}) {
		return model.APIResponse{}, api.NewDuplicateErr(err.Error())
	} else if errors.As(err, &services.ValidationErr{}) {
		return model.APIResponse{}, api.NewBadFormatErr(err.Error())
	}
	return model.APIResponse{}, err
}
