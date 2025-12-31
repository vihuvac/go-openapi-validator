# Changelog

All notable changes to this project will be documented in this file. The format
is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this
project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2025-12-31

**Initial release of Go OpenAPI Validator** – a high-performance, framework-agnostic OpenAPI v3 validator for Go.

### Added

- **Framework Agnostic Middleware**: Core validation engine that works seamlessly with `net/http`, [Gorilla Mux](https://github.com/gorilla/mux), and [Gin-gonic](https://github.com/gin-gonic/gin).
- **OpenAPI v3 Validation**: Full support for validating request bodies, query parameters, and headers against OpenAPI v3 specifications using [kin-openapi](https://github.com/getkin/kin-openapi).
- **Response Validation**: Optional validation of outgoing responses to ensure API compliance.
- **Embedded Swagger UI**: Built-in Swagger UI integration served directly from the application using Go 1.16+ `embed` functionality.
- **Flexible Configuration**: Implementation of the Functional Options pattern for easy and readable configuration (e.g., `WithValidateRequests`, `WithSwaggerUIPath`).
- **Customizable Error Handling**: Support for custom `ErrorEncoder` to allow users to define their own error response formats.
- **Zero Testing Dependencies**: Comprehensive test suite implemented using only the Go standard library for a lean and reliable codebase.
- **Interactive Documentation**: Accessible and zero-config API documentation served at `/docs` (configurable).
- **Multi-Framework Examples**: Detailed integration examples for standard library, Gorilla Mux, and Gin.
- **Repository tooling**: Added issue templates and GitHub workflows (CI, release/publish).
