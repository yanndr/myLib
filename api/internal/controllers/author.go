package controllers

import (
	"api/api"
	"api/internal/services"
	"api/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
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

// Get is the endpoint action for the GET method to retrieve an author.
func (c *AuthorController) Get(r *http.Request) (model.APIResponse, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return model.APIResponse{}, api.NewBadFormatErr(err.Error())
	}

	author, err := c.AuthorService.GetById(r.Context(), id)
	if err != nil {
		return handleServiceError(err)
	}

	return model.NewContentResponse(author), nil
}

// Delete is the endpoint action for the DELETE method for deleting an author.
func (c *AuthorController) Delete(r *http.Request) (model.APIResponse, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return model.APIResponse{}, api.NewBadFormatErr(err.Error())
	}

	err = c.AuthorService.Delete(r.Context(), id)
	if err != nil {
		return handleServiceError(err)
	}

	return model.NewEmptyResponse(), nil
}

// GetAll is the endpoint action for the GET method to retrieve the list of authors.
func (c *AuthorController) GetAll(r *http.Request) (model.APIResponse, error) {
	lastname := r.URL.Query().Get("lastname")

	authors, err := c.AuthorService.GetAll(r.Context(), lastname)
	if err != nil {
		return handleServiceError(err)
	}

	return model.NewContentResponse(authors), nil
}

func handleServiceError(err error) (model.APIResponse, error) {
	if errors.As(err, &services.DuplicateErr{}) {
		return model.APIResponse{}, api.NewDuplicateErr(err.Error())
	} else if errors.As(err, &services.ValidationErr{}) {
		return model.APIResponse{}, api.NewBadFormatErr(err.Error())
	} else if errors.As(err, &services.NotFoundErr{}) {
		return model.APIResponse{}, api.NewNotFoundErr(err.Error())
	}
	return model.APIResponse{}, err
}
