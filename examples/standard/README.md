# Standard Library Example

This example demonstrates how to use `go-openapi-validator` with the Go standard library's `http.ServeMux`.

## Features

- **Minimal Dependencies**: Shows how to use the validator with standard `net/http`.
- **Legacy Router Support**: Uses the `legacy` router from `kin-openapi` (required for `ServeMux`).
- **Post Request Validation**: Includes an example of validating JSON bodies in POST requests.

## Usage

### 1. Run the Server

```bash
go run main.go
```

### 2. Test Endpoints

- **Hello Endpoint**:
  ```bash
  curl -X POST http://localhost:8081/hello -H "Content-Type: application/json" -d '{"name": "Developer"}'
  ```
- **Validation Error Example** (Missing name):
  ```bash
  curl -X POST http://localhost:8081/hello -H "Content-Type: application/json" -d '{}'
  ```
- **Swagger UI**: Visit `http://localhost:8081/docs` in your browser.

## Code Overview

For `net/http`, we manually configure the router to use the `legacy` implementation:

```go
v, _ := validator.New("openapi.yaml")
router, _ := legacy.NewRouter(v.Swagger)
validator.WithRouter(router)(v.Options)

mux := http.NewServeMux()
// ... register routes ...

handler := v.Middleware(mux)
http.ListenAndServe(":8081", handler)
```
