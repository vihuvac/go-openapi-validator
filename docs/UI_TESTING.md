# UI Testing Guide

This guide explains how to verify the integrated Swagger UI and ensure your OpenAPI specification is rendering correctly.

## Verifying Swagger UI Access

By default, Swagger UI is served at `/docs`. You can verify it by running your server and navigating to:

`http://localhost:8080/docs`

You should see the Swagger UI dashboard with your specification loaded.

## Testing Your Specification

1. **Check for Errors**: If the UI doesn't load or shows an error, check the browser console for JavaScript errors.
2. **Verify Endpoints**: Ensure all paths and methods defined in your `openapi.yaml` are visible in the UI.
3. **Try it Out**: Use the "Try it out" button to send requests directly from the UI. These requests will go through the validation middleware if it's applied to your mux.

## Changing the UI Path

If you want to serve Swagger UI at a different path, use the `WithSwaggerUIPath` option:

```go
v, _ := validator.New("openapi.yaml", validator.WithSwaggerUIPath("/api-docs"))
```

Then access it at `http://localhost:8080/api-docs`.

## Validating the Spec File

The UI fetches the specification from `<SwaggerUIPath>/openapi.json`. You can verify this endpoint directly with `curl`:

```bash
curl http://localhost:8080/docs/openapi.json
```
