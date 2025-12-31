package openapi_validator

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const testSpec = `
openapi: 3.0.0
info:
  title: Test API
  version: 1.0.0
paths:
  /test:
    post:
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              required: [name]
              properties:
                name: {type: string}
      responses:
        '200':
          description: OK
          content:
            application/json:
              schema:
                type: object
                properties:
                  result: {type: string}
`

func TestValidator_Middleware_RequestValidation(t *testing.T) {
	tmpSpec := "test_spec.yaml"
	err := os.WriteFile(tmpSpec, []byte(testSpec), 0644)
	if err != nil {
		t.Fatalf("failed to write temp spec: %v", err)
	}
	defer os.Remove(tmpSpec)

	v, err := New(tmpSpec)
	if err != nil {
		t.Fatalf("failed to create validator: %v", err)
	}

	handler := v.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result":"ok"}`))
	}))

	t.Run("Valid Request", func(t *testing.T) {
		// Arrange
		body := `{"name":"test"}`
		req := httptest.NewRequest("POST", "/test", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Act
		handler.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}
	})

	t.Run("Invalid Request - Missing Field", func(t *testing.T) {
		// Arrange
		body := `{"wrong":"field"}`
		req := httptest.NewRequest("POST", "/test", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Act
		handler.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", w.Code)
		}

		var resp ValidationError
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("failed to unmarshal response: %v", err)
		}

		if resp.Message != "Validation Failed" {
			t.Errorf("expected Message 'Validation Failed', got %s", resp.Message)
		}
	})
}

func TestNew(t *testing.T) {
	t.Run("Invalid File Path", func(t *testing.T) {
		// Arrange & Act
		_, err := New("non_existent.yaml")

		// Assert
		if err == nil {
			t.Error("expected error for non-existent file")
		}
	})

	t.Run("Invalid Spec Content", func(t *testing.T) {
		// Arrange
		tmpSpec := "invalid_spec.yaml"
		os.WriteFile(tmpSpec, []byte("invalid yaml content"), 0644)
		defer os.Remove(tmpSpec)

		// Act
		_, err := New(tmpSpec)

		// Assert
		if err == nil {
			t.Error("expected error for invalid spec")
		}
	})
}

func TestValidator_Middleware_ResponseValidation(t *testing.T) {
	tmpSpec := "test_spec_resp.yaml"
	os.WriteFile(tmpSpec, []byte(testSpec), 0644)
	defer os.Remove(tmpSpec)

	v, _ := New(tmpSpec, WithValidateResponses(true))

	t.Run("Valid Response", func(t *testing.T) {
		// Arrange
		handler := v.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"result":"ok"}`))
		}))

		req := httptest.NewRequest("POST", "/test", bytes.NewBufferString(`{"name":"test"}`))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Act
		handler.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}
	})

	t.Run("Invalid Response Body", func(t *testing.T) {
		// Arrange
		handler := v.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"wrong":"field"}`))
		}))

		req := httptest.NewRequest("POST", "/test", bytes.NewBufferString(`{"name":"test"}`))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Act
		// This should log a warning but still return 200 since it's response validation
		handler.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}
	})

	t.Run("Route Not Found Pass Through", func(t *testing.T) {
		// Arrange
		handler := v.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))

		req := httptest.NewRequest("GET", "/not-exists", nil)
		w := httptest.NewRecorder()

		// Act
		handler.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", w.Code)
		}
	})
}
