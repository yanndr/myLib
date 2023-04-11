package model

import "net/http"

const (
	SuccessStatus = "Success"
	ErrorStatus   = "Error"
)

// APIResponse represents the default response to an API request.
type APIResponse struct {
	StatusCode   int    `json:"status_code"`
	Status       string `json:"status"`
	Content      any    `json:"content,omitempty"`
	Location     string `json:"location,omitempty"`
	Error        string `json:"error,omitempty"`
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorDetails string `json:"error_details,omitempty"`
}

// NewContentResponse creates a new APIResponse with StatusCode "Status Ok" and Status "Success".
// It also sets the content from the parameter.
func NewContentResponse(content any) APIResponse {
	return APIResponse{
		StatusCode: http.StatusOK,
		Status:     SuccessStatus,
		Content:    content,
	}
}

// NewErrorResponse creates a new APIResponse with Status "ErrorStatus"
// and sets the StatusCode, Error and ErrorDetails from the parameters.
func NewErrorResponse(statusCode int, err, errorCode, details string) APIResponse {
	return APIResponse{
		StatusCode:   statusCode,
		Status:       ErrorStatus,
		Error:        err,
		ErrorDetails: details,
		ErrorCode:    errorCode,
	}
}

// NewCreatedResponse creates a new APIResponse with StatusCode "StatusCreated" and Status "Success".
// Its also sets the location from the parameter
func NewCreatedResponse(location string) APIResponse {
	return APIResponse{
		StatusCode: http.StatusCreated,
		Status:     SuccessStatus,
		Location:   location,
	}
}
