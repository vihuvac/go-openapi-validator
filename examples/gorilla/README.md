# Gorilla Mux Example

This example demonstrates how to integrate `go-openapi-validator` with the popular [Gorilla Mux](https://github.com/gorilla/mux) router.

## Features

- **Native Router Support**: Leverages `kin-openapi`'s built-in Gorilla Mux router for path matching.
- **Standard Middleware**: Uses the standard `http.Handler` middleware pattern.
- **Spec Validation**: Full request and response validation against `openapi.yaml`.

## Usage

### 1. Run the Server

```bash
go run main.go
```

### 2. Test Endpoints

- **Liveness Check**: `curl http://localhost:8081/health/liveness`
- **Readiness Check**: `curl http://localhost:8081/health/readiness`
- **Swagger UI**: Visit `http://localhost:8081/docs` in your browser.

## Code Overview

Initialization is straightforward as `Gorilla Mux` is the default router used by the validator:

```go
v, _ := validator.New("openapi.yaml")
r := mux.NewRouter()

// Define routes
r.HandleFunc("/health/liveness", handleLiveness).Methods("GET")

// Apply middleware
handler := v.Middleware(r)
http.ListenAndServe(":8081", handler)
```
