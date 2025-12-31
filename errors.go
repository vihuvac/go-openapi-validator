package openapi_validator

import (
	"encoding/json"
	"net/http"
)

// ValidationError represents a failed validation attempt.
// It includes a top-level message and an optional list of detailed error strings.
type ValidationError struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}

// Error implements the error interface for ValidationError.
func (e *ValidationError) Error() string {
	return e.Message
}

// DefaultErrorEncoder is a built-in implementation of ErrorEncoder.
// It sends a JSON response with a 400 Bad Request status code.
func DefaultErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	resp := ValidationError{
		Message: "Validation Failed",
		Errors:  []string{err.Error()},
	}

	json.NewEncoder(w).Encode(resp)
}
