# Gin Framework Example (Handler Mapping)

This advanced example demonstrates how to use `operationId` and custom OpenAPI extensions (`x-eov-operation-handler`) to automatically register routes in the [Gin-gonic](https://github.com/gin-gonic/gin) framework.

## Key Concepts

- **Automated Routing**: Routes are dynamically registered by iterating over the OpenAPI specification.
- **Spec-Driven Development**: Define your API structure in `openapi.yaml` and let the code handle the plumbing.
- **Handler Mapping**: Uses a simple map to link `operationId` from the spec to Go controller methods.

## Usage

### 1. Run the Server

```bash
go run *.go
```

### 2. Test Endpoints

- **Liveness Check**: `curl http://localhost:8081/health/liveness`
- **Readiness Check**: `curl http://localhost:8081/health/readiness`
- **Swagger UI**: Visit `http://localhost:8081/docs` in your browser.

## How it Works

The application iterates through `v.Swagger.Paths` and matches the `OperationID` with a predefined map of handlers:

```go
handlers := map[string]gin.HandlerFunc{
  "CheckLiveness": healthController.CheckLiveness,
}

for path, pathItem := range v.Swagger.Paths.Map() {
  for method, operation := range pathItem.Operations() {
    if handler, ok := handlers[operation.OperationID]; ok {
      r.Handle(method, path, handler)
    }
  }
}
```
