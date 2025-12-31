# go-openapi-validator Examples

This directory contains several examples demonstrating how to integrate `go-openapi-validator` with different Go web routers and frameworks. These examples cover various use cases, from basic manual route registration to advanced automated handler mapping.

## Documentation Index

| Example | Router / Framework | Description |
| :--- | :--- | :--- |
| [**Gin (Regular)**](./gin) | [Gin-gonic](https://github.com/gin-gonic/gin) | Manual route registration with OpenAPI validation middleware. |
| [**Gin (Handler Mapping)**](./gin-handler-mapping) | [Gin-gonic](https://github.com/gin-gonic/gin) | Automated route registration using `operationId` and custom extensions. |
| [**Gorilla Mux**](./gorilla) | [Gorilla Mux](https://github.com/gorilla/mux) | Standard integration with the popular Gorilla Mux router. |
| [**Standard Library**](./standard) | `net/http` | Integration with the standard library `http.ServeMux`. |

## Key Features Demonstrated

- **Request Validation**: Ensuring incoming requests match the OpenAPI specification.
- **Response Validation**: Validating outgoing responses against the spec (useful for development).
- **Swagger UI Integration**: Serving a built-in Swagger UI with live spec validation.
- **Automated Routing**: Reducing boilerplate by mapping OpenAPI `operationId` to Go handlers.

## Getting Started

To run any of the examples, navigate to the respective directory and use the standard `go run` command:

```bash
cd examples/gin
go run main.go
```

Once the server is running, you can access the Swagger UI at `http://localhost:8081/docs`.

---

**[Back to main project](../../README.md)**
