package openapi_validator

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestValidator_HandleSwaggerUI(t *testing.T) {
	tmpSpec := "test_spec_swagger_ui.yaml"
	err := os.WriteFile(tmpSpec, []byte(testSpec), 0644)
	if err != nil {
		t.Fatalf("failed to write temp spec: %v", err)
	}
	defer os.Remove(tmpSpec)

	v, err := New(tmpSpec)
	if err != nil {
		t.Fatalf("failed to create validator: %v", err)
	}

	mux := http.NewServeMux()
	v.HandleSwaggerUI(mux)

	t.Run("Serve index HTML", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest("GET", "/docs/", nil)
		w := httptest.NewRecorder()

		// Act
		mux.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		if w.Header().Get("Content-Type") != "text/html" {
			t.Errorf("expected Content-Type text/html, got %s", w.Header().Get("Content-Type"))
		}

		body := w.Body.String()
		if !strings.Contains(body, "<title>Swagger UI</title>") {
			t.Error("body does not contain expected title")
		}

		if !strings.Contains(body, "https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js") {
			t.Error("body does not contain swagger bundle js")
		}

		if !strings.Contains(body, "styles.css") {
			t.Error("body does not contain styles.css")
		}

		if !strings.Contains(body, "main.js") {
			t.Error("body does not contain main.js")
		}
	})

	t.Run("Serve openapi.json", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest("GET", "/docs/openapi.json", nil)
		w := httptest.NewRecorder()

		// Act
		mux.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		if w.Header().Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type application/json, got %s", w.Header().Get("Content-Type"))
		}
	})

	t.Run("Serve styles.css", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest("GET", "/docs/styles.css", nil)
		w := httptest.NewRecorder()

		// Act
		mux.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		if w.Header().Get("Content-Type") != "text/css" {
			t.Errorf("expected Content-Type text/css, got %s", w.Header().Get("Content-Type"))
		}

		if !strings.Contains(w.Body.String(), ".swagger-validator") {
			t.Error("body does not contain expected css class")
		}
	})

	t.Run("Serve main.js", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest("GET", "/docs/main.js", nil)
		w := httptest.NewRecorder()

		// Act
		mux.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		if w.Header().Get("Content-Type") != "application/javascript" {
			t.Errorf("expected Content-Type application/javascript, got %s", w.Header().Get("Content-Type"))
		}

		if !strings.Contains(w.Body.String(), "SwaggerUIBundle") {
			t.Error("body does not contain SwaggerUIBundle")
		}
	})

	t.Run("Serve 404 for invalid path", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest("GET", "/docs/not-found", nil)
		w := httptest.NewRecorder()

		// Act
		mux.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", w.Code)
		}
	})

	t.Run("Serve index.html explicitly", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest("GET", "/docs/index.html", nil)
		w := httptest.NewRecorder()

		// Act
		mux.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		if w.Header().Get("Content-Type") != "text/html" {
			t.Errorf("expected Content-Type text/html, got %s", w.Header().Get("Content-Type"))
		}
	})
}

func TestResponseWriter_WriteDefaultStatus(t *testing.T) {
	// Arrange
	w := httptest.NewRecorder()
	rw := &responseWriter{
		ResponseWriter: w,
		header:         w.Header(),
	}

	// Act
	rw.Write([]byte("test"))

	// Assert
	if rw.status != http.StatusOK {
		t.Errorf("expected status 200, got %d", rw.status)
	}
}

func TestValidator_HandleSwaggerUI_CustomPath(t *testing.T) {
	tmpSpec := "test_spec_custom_ui.yaml"
	err := os.WriteFile(tmpSpec, []byte(testSpec), 0644)
	if err != nil {
		t.Fatalf("failed to write temp spec: %v", err)
	}
	defer os.Remove(tmpSpec)

	customPath := "/api-docs"
	v, err := New(tmpSpec, WithSwaggerUIPath(customPath))
	if err != nil {
		t.Fatalf("failed to create validator: %v", err)
	}

	mux := http.NewServeMux()
	v.HandleSwaggerUI(mux)

	t.Run("Serve index HTML on custom path", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest("GET", customPath+"/", nil)
		w := httptest.NewRecorder()

		// Act
		mux.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		if !strings.Contains(w.Body.String(), customPath+"/openapi.json") {
			t.Error("body does not contain custom openapi.json path")
		}
	})

	t.Run("Serve assets on custom path", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest("GET", customPath+"/styles.css", nil)
		w := httptest.NewRecorder()

		// Act
		mux.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}
	})
}

func TestValidator_HandleSwaggerUI_NoTrailingSlash(t *testing.T) {
	// Arrange
	tmpSpec := "test_spec_no_slash.yaml"
	os.WriteFile(tmpSpec, []byte(testSpec), 0644)
	defer os.Remove(tmpSpec)

	v, _ := New(tmpSpec, WithSwaggerUIPath("/docs"))
	mux := http.NewServeMux()
	v.HandleSwaggerUI(mux)

	req := httptest.NewRequest("GET", "/docs/", nil)
	w := httptest.NewRecorder()

	// Act
	mux.ServeHTTP(w, req)

	// Assert
	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}
