# Gin Framework Example (Regular)

This example demonstrates how to integrate `go-openapi-validator` with the [Gin-gonic](https://github.com/gin-gonic/gin) framework using manual route registration.

## Features

- **Middleware Integration**: Seamlessly validates requests and responses using Gin's middleware pattern.
- **Manual Routing**: Standard Gin route definitions (`r.GET`, `r.POST`, etc.).
- **Automatic Documentation**: Serves Swagger UI at `/docs`.

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

The validator is initialized and then used as a middleware wrapper for the Gin engine:

```go
v, _ := validator.New("openapi.yaml")
r := gin.Default()

// Register routes manually
r.GET("/health/liveness", handleLiveness)

// Wrap and serve
handler := v.Middleware(r)
http.ListenAndServe(":8081", handler)
```
