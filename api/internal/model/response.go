package model

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
