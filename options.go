package openapi_validator

import (
	"net/http"

	"github.com/getkin/kin-openapi/routers"
)

// Option is a function type used to configure the Options struct.
type Option func(*Options)

// Options contains the configuration for the OpenAPI validator.
type Options struct {
	// ValidateRequests specifies whether incoming requests should be validated against the spec.
	ValidateRequests bool
	// ValidateResponses specifies whether outgoing responses should be validated against the spec.
	ValidateResponses bool
	// SwaggerUIPath is the URL path where Swagger UI will be served.
	SwaggerUIPath string
	// ErrorEncoder is used to format and send validation error responses.
	ErrorEncoder ErrorEncoder
	// Router is used for matching requests to OpenAPI paths.
	Router routers.Router
}

// ErrorEncoder is a function type used to encode validation errors into an HTTP response.
type ErrorEncoder func(w http.ResponseWriter, r *http.Request, err error)

// DefaultOptions returns the default configuration for the validator.
func DefaultOptions() *Options {
	return &Options{
		ValidateRequests:  true,
		ValidateResponses: false,
		SwaggerUIPath:     "/docs",
		ErrorEncoder:      DefaultErrorEncoder,
	}
}

// WithValidateRequests returns an Option that enables or disables request validation.
func WithValidateRequests(validate bool) Option {
	return func(o *Options) {
		o.ValidateRequests = validate
	}
}

// WithValidateResponses returns an Option that enables or disables response validation.
func WithValidateResponses(validate bool) Option {
	return func(o *Options) {
		o.ValidateResponses = validate
	}
}

// WithSwaggerUIPath returns an Option that sets the URL path for the Swagger UI.
func WithSwaggerUIPath(path string) Option {
	return func(o *Options) {
		o.SwaggerUIPath = path
	}
}

// WithErrorEncoder returns an Option that sets a custom error encoder.
func WithErrorEncoder(encoder ErrorEncoder) Option {
	return func(o *Options) {
		o.ErrorEncoder = encoder
	}
}

// WithRouter returns an Option that sets a custom router.
func WithRouter(router routers.Router) Option {
	return func(o *Options) {
		o.Router = router
	}
}
