# Development Guide

Welcome to the `go-openapi-validator` development guide! This document explains how the package is structured and how you can contribute.

## Architecture

The package follows a simple, middleware-based architecture:

- **`Validator`**: The core struct that holds the parsed OpenAPI spec and the router.
- **`Middleware`**: A standard `net/http` middleware that intercepts requests and responses for validation.
- **`openapi3filter`**: We use the excellent `kin-openapi` library for the actual validation logic.
- **`Swagger UI`**: Embedded assets from the official Swagger UI project are served via a dedicated handler.

## Project Structure

- `validator.go`: Core middleware and validator logic.
- `options.go`: Configuration options for the validator.
- `errors.go`: Custom error handling and JSON encoding.
- `swagger.go`: Swagger UI integration and static file serving.
- `swagger-ui/`: Directory containing the embedded Swagger UI assets.

## Running Tests

We value high test coverage. You can run the tests using the standard Go toolchain:

```bash
go test ./...
```

For more detailed output and coverage:

```bash
go test -v -cover ./...
```

## Contributing

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Write tests for your changes.
4. Ensure all tests pass.
5. Submit a pull request.
