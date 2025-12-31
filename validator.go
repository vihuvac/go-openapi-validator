package openapi_validator

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
)

// Validator is the core component that manages OpenAPI validation and Swagger UI.
type Validator struct {
	// Options holds the configuration for the validator.
	Options *Options
	// Swagger is the parsed OpenAPI 3 specification.
	Swagger *openapi3.T
}

// New creates a new Validator instance from an OpenAPI spec file and optional configuration.
// It parses and validates the spec, and initializes the router.
func New(specPath string, opts ...Option) (*Validator, error) {
	ctx := context.Background()
	loader := openapi3.NewLoader()
	swagger, err := loader.LoadFromFile(specPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load spec: %w", err)
	}

	if err := swagger.Validate(ctx); err != nil {
		return nil, fmt.Errorf("invalid spec: %w", err)
	}

	options := DefaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	// Default to gorillamux if no router provided
	if options.Router == nil {
		router, err := gorillamux.NewRouter(swagger)
		if err != nil {
			return nil, fmt.Errorf("failed to create default router: %w", err)
		}
		options.Router = router
	}

	return &Validator{
		Options: options,
		Swagger: swagger,
	}, nil
}

// Middleware returns an http.Handler that validates incoming requests and/or outgoing responses.
func (v *Validator) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip validation for Swagger UI
		if strings.HasPrefix(r.URL.Path, v.Options.SwaggerUIPath) {
			next.ServeHTTP(w, r)
			return
		}

		// Find route
		route, pathParams, err := v.Options.Router.FindRoute(r)
		if err != nil {
			// If route not found, we can decide to either block or pass through.
			// express-openapi-validator usually lets it pass if not explicitly defined?
			// but for a strict validator, it's better to log or return error.
			// Let's pass through for now, as not all routes might be in OpenAPI spec.
			next.ServeHTTP(w, r)
			return
		}

		// Validate Request
		if v.Options.ValidateRequests {
			requestValidationInput := &openapi3filter.RequestValidationInput{
				Request:    r,
				PathParams: pathParams,
				Route:      route,
			}
			if err := openapi3filter.ValidateRequest(context.Background(), requestValidationInput); err != nil {
				v.Options.ErrorEncoder(w, r, err)
				return
			}
		}

		// Response validation
		if v.Options.ValidateResponses {
			rw := &responseWriter{
				ResponseWriter: w,
				header:         w.Header(),
			}
			next.ServeHTTP(rw, r)

			// After handler. Check if we should validate
			responseValidationInput := &openapi3filter.ResponseValidationInput{
				RequestValidationInput: &openapi3filter.RequestValidationInput{
					Request:    r,
					PathParams: pathParams,
					Route:      route,
				},
				Status: rw.status,
				Header: rw.header,
			}

			if rw.body != nil {
				responseValidationInput.SetBodyBytes(rw.body)
			}

			if err := openapi3filter.ValidateResponse(context.Background(), responseValidationInput); err != nil {
				// NOTE: We already sent the response to the user.
				// Response validation is mostly for development/logging.
				// We could log it here.
				fmt.Printf("Response validation error: %v\n", err)
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
	body   []byte
	header http.Header
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if rw.status == 0 {
		rw.status = http.StatusOK
	}
	rw.body = append(rw.body, b...)
	return rw.ResponseWriter.Write(b)
}
