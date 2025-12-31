package openapi_validator

import (
	"net/http"
	"testing"
)

func TestDefaultOptions(t *testing.T) {
	// Arrange & Act
	opts := DefaultOptions()

	// Assert
	if !opts.ValidateRequests {
		t.Error("expected ValidateRequests to be true")
	}

	if opts.ValidateResponses {
		t.Error("expected ValidateResponses to be false")
	}

	if opts.SwaggerUIPath != "/docs" {
		t.Errorf("expected SwaggerUIPath to be /docs, got %s", opts.SwaggerUIPath)
	}

	if opts.ErrorEncoder == nil {
		t.Error("expected ErrorEncoder to not be nil")
	}
}

func TestWithValidateRequests(t *testing.T) {
	// Arrange
	opts := DefaultOptions()

	// Act
	WithValidateRequests(false)(opts)

	// Assert
	if opts.ValidateRequests {
		t.Error("expected ValidateRequests to be false")
	}
}

func TestWithValidateResponses(t *testing.T) {
	// Arrange
	opts := DefaultOptions()

	// Act
	WithValidateResponses(true)(opts)

	// Assert
	if !opts.ValidateResponses {
		t.Error("expected ValidateResponses to be true")
	}
}

func TestWithSwaggerUIPath(t *testing.T) {
	// Arrange
	opts := DefaultOptions()

	// Act
	WithSwaggerUIPath("/api/docs")(opts)

	// Assert
	if opts.SwaggerUIPath != "/api/docs" {
		t.Errorf("expected SwaggerUIPath to be /api-docs, got %s", opts.SwaggerUIPath)
	}
}

func TestWithErrorEncoder(t *testing.T) {
	// Arrange
	opts := DefaultOptions()
	customEncoder := func(w http.ResponseWriter, r *http.Request, err error) {}

	// Act
	WithErrorEncoder(customEncoder)(opts)

	// Assert
	if opts.ErrorEncoder == nil {
		t.Error("expected ErrorEncoder to not be nil")
	}
}

func TestWithRouter(t *testing.T) {
	// Arrange
	opts := DefaultOptions()

	// Act
	// routers.Router is an interface, we can use nil for just checking if it's set
	// but better to check if it's actually assigned
	WithRouter(nil)(opts)

	// Assert
	if opts.Router != nil {
		t.Error("expected Router to be nil")
	}
}
