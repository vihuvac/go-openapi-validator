package openapi_validator

import (
	"embed"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
)

//go:embed swagger-ui/*
var swaggerUIFS embed.FS

var (
	swaggerUITemplate = template.Must(template.ParseFS(swaggerUIFS, "swagger-ui/index.html"))
)

// Registrar is an interface that matches both http.ServeMux and gorilla/mux.Router.
type Registrar interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

// SwaggerUIHandler returns an http.Handler that serves the Swagger UI and the OpenAPI spec.
func (v *Validator) SwaggerUIHandler() http.Handler {
	path := v.Options.SwaggerUIPath
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve the current file from the swagger-ui directory
		relPath := strings.TrimPrefix(r.URL.Path, path)

		// Serve the spec file if requested
		if relPath == "openapi.json" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(v.Swagger)
			return
		}

		// Serve static assets from the embedded filesystem
		if relPath == "styles.css" || relPath == "main.js" {
			content, err := swaggerUIFS.ReadFile("swagger-ui/" + relPath)
			if err != nil {
				http.Error(w, "Asset not found", http.StatusNotFound)
				return
			}
			if strings.HasSuffix(relPath, ".css") {
				w.Header().Set("Content-Type", "text/css")
			} else {
				w.Header().Set("Content-Type", "application/javascript")
			}
			w.Write(content)
			return
		}

		// Serve the index HTML for the base path or index.html explicitly
		if relPath == "" || relPath == "index.html" {
			w.Header().Set("Content-Type", "text/html")
			data := struct {
				SpecURL string
			}{
				SpecURL: v.Options.SwaggerUIPath + "/openapi.json",
			}
			if err := swaggerUITemplate.Execute(w, data); err != nil {
				http.Error(w, "Failed to render Swagger UI", http.StatusInternalServerError)
			}
			return
		}

		// Otherwise, return 404
		http.NotFound(w, r)
	})
}

// HandleSwaggerUI registers the necessary routes to serve the Swagger UI and the OpenAPI spec.
func (v *Validator) HandleSwaggerUI(mux Registrar) {
	path := v.Options.SwaggerUIPath
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	handler := v.SwaggerUIHandler()
	mux.HandleFunc(path, handler.ServeHTTP)
}
