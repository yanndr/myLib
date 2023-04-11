package endpoints

import (
	"api/internal/api"
	"api/internal/model"
	"api/internal/services"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	AuthorsPath = "/authors"
	BooksPath   = "/books"
	apiV1       = "/v1"
)

var NotImplementedErr = fmt.Errorf("not implemented")

// EndpointHandler represents a request method that returns a model.APIResponse.
type EndpointHandler func(r *http.Request) (model.APIResponse, error)

// Route represents a specific route of a server. It contains all the EndpointHandler mapped
// to their http.methods and all the sub routes
type Route struct {
	Pattern   string
	Actions   map[string]EndpointHandler
	SubRoutes []*Route
}

// NewV1Route creates the V1 route of the API
func NewV1Route(apiVersion string, authSvc services.AuthorService) *Route {
	return &Route{
		Pattern: apiV1,
		Actions: map[string]EndpointHandler{
			http.MethodGet: func(r *http.Request) (model.APIResponse, error) {
				return model.NewContentResponse(&model.APIInformation{APIVersion: apiVersion}), nil
			},
		},
		SubRoutes: []*Route{
			newAuthorsEndpoint(apiV1, authSvc),
		},
	}
}

// Handle responds to an HTTP request,  executes the action EndpointHandler
// and writes the result to the http.ResponseWriter
func Handle(action EndpointHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response, err := action(r)
		if err != nil {
			response = getErrorResponse(err)
		}
		w.WriteHeader(response.StatusCode)
		_ = json.NewEncoder(w).Encode(response)
	}
}

func notImplementedHandler(_ *http.Request) (model.APIResponse, error) {
	return model.APIResponse{}, NotImplementedErr
}

func getErrorResponse(err error) model.APIResponse {
	var statusCode int
	var errorCode string
	var errMessage string
	var details string
	var statusError api.StatusErr
	if errors.As(err, &statusError) {
		statusCode = statusError.StatusCode
		errorCode = statusError.ErrorCode
		errMessage = statusError.ErrorMessage
		details = statusError.ErrorDetails
	} else {
		statusCode = http.StatusInternalServerError
		errorCode = api.UnexpectedErrorCode
		errMessage = api.UnexpectedErrorMessage
		details = err.Error()
	}

	return model.NewErrorResponse(statusCode, errMessage, errorCode, details)
}
