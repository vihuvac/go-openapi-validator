package openapi_validator

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestValidationError(t *testing.T) {
	// Arrange
	err := &ValidationError{
		Message: "test error",
		Errors:  []string{"detail 1", "detail 2"},
	}

	// Act
	got := err.Error()

	// Assert
	if got != "test error" {
		t.Errorf("expected test error, got %s", got)
	}
}

func TestDefaultErrorEncoder(t *testing.T) {
	// Arrange
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	err := errors.New("test validation error")

	// Act
	DefaultErrorEncoder(w, r, err)

	// Assert
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}

	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", w.Header().Get("Content-Type"))
	}

	var resp ValidationError
	if errUnmarshal := json.Unmarshal(w.Body.Bytes(), &resp); errUnmarshal != nil {
		t.Fatalf("failed to unmarshal response: %v", errUnmarshal)
	}

	if resp.Message != "Validation Failed" {
		t.Errorf("expected Message Validation Failed, got %s", resp.Message)
	}

	if len(resp.Errors) == 0 || !strings.Contains(resp.Errors[0], "test validation error") {
		t.Errorf("expected error details to contain 'test validation error', got %v", resp.Errors)
	}
}
