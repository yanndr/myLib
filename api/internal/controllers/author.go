package controllers

import (
	"api/api"
	"api/internal/services"
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
func (c *AuthorController) Create(r *http.Request) (api.Response, error) {
	var author api.CreateUpdateAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		return api.Response{}, api.NewBadFormatErr(err.Error())
	}
	if author == (api.CreateUpdateAuthorRequest{}) {
		return api.Response{}, api.NewBadFormatErr("the author is empty")
	}

	id, err := c.AuthorService.Create(r.Context(), author.AuthorBase)
	if err != nil {
		return handleServiceError(err)
	}

	location := fmt.Sprintf("%v/%v", c.BasePath, id)
	return api.NewCreatedResponse(location), nil
}

// Get is the endpoint action for the GET method to retrieve an author.
func (c *AuthorController) Get(r *http.Request) (api.Response, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return api.Response{}, api.NewBadFormatErr(err.Error())
	}

	author, err := c.AuthorService.GetById(r.Context(), id)
	if err != nil {
		return handleServiceError(err)
	}

	return api.NewContentResponse(author), nil
}

// Delete is the endpoint action for the DELETE method for deleting an author.
func (c *AuthorController) Delete(r *http.Request) (api.Response, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return api.Response{}, api.NewBadFormatErr(err.Error())
	}

	err = c.AuthorService.Delete(r.Context(), id)
	if err != nil {
		return handleServiceError(err)
	}

	return api.NewEmptyResponse(), nil
}

// GetAll is the endpoint action for the GET method to retrieve the list of authors.
func (c *AuthorController) GetAll(r *http.Request) (api.Response, error) {
	lastname := r.URL.Query().Get("lastname")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		//log
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		//log
	}

	authors, err := c.AuthorService.GetAll(r.Context(), lastname, limit, offset)
	if err != nil {
		return handleServiceError(err)
	}

	return api.NewContentResponse(authors), nil
}

// Update is the endpoint action for the PUT method for updating an author.
func (c *AuthorController) Update(r *http.Request) (api.Response, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return api.Response{}, api.NewBadFormatErr(err.Error())
	}

	var author api.CreateUpdateAuthorRequest
	if err = json.NewDecoder(r.Body).Decode(&author); err != nil {
		return api.Response{}, api.NewBadFormatErr(err.Error())
	}

	ctx := r.Context()
	eTag := r.Header.Get("If-Match")
	if eTag != "" {
		existing, err := c.AuthorService.GetById(ctx, id)
		if err != nil {
			return api.Response{}, err
		}
		existingEtag, err := existing.Serialize()

		if existingEtag != eTag {
			return api.Response{}, api.NewPreconditionFailError()
		}
	}

	err = c.AuthorService.Update(ctx, id, author.AuthorBase)
	if err != nil {
		return handleServiceError(err)
	}

	return api.NewEmptyResponse(), nil
}

// PartialUpdate is the endpoint action for the PATCH method for partially updating an author.
func (c *AuthorController) PartialUpdate(r *http.Request) (api.Response, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return api.Response{}, api.NewBadFormatErr(err.Error())
	}

	var patchAuthorRequest api.PatchAuthorRequest
	if err = json.NewDecoder(r.Body).Decode(&patchAuthorRequest); err != nil {
		return api.Response{}, api.NewBadFormatErr(err.Error())
	}

	err = c.AuthorService.PartialUpdate(r.Context(), id, patchAuthorRequest)
	if err != nil {
		return handleServiceError(err)
	}

	return api.NewEmptyResponse(), nil
}

func handleServiceError(err error) (api.Response, error) {
	if errors.As(err, &services.DuplicateErr{}) {
		return api.Response{}, api.NewDuplicateErr(err.Error())
	} else if errors.As(err, &services.ValidationErr{}) {
		return api.Response{}, api.NewBadFormatErr(err.Error())
	} else if errors.As(err, &services.NotFoundErr{}) {
		return api.Response{}, api.NewNotFoundErr(err.Error())
	}
	return api.Response{}, err
}
