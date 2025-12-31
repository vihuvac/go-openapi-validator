package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/vihuvac/go-openapi-validator"
)

func main() {
	// Gin can work with the default gorillamux router for matching.
	v, err := validator.New(
		"openapi.yaml",
		validator.WithValidateRequests(true),
		validator.WithValidateResponses(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	// Register Swagger UI using SwaggerUIHandler for Gin-gonic.
	r.GET(v.Options.SwaggerUIPath+"/*any", gin.WrapH(v.SwaggerUIHandler()))

	// Define handler mapping
	healthController := &HealthController{}
	handlers := map[string]gin.HandlerFunc{
		"CheckLiveness":  healthController.CheckLiveness,
		"CheckReadiness": healthController.CheckReadiness,
	}

	// Automate route registration based on OpenAPI spec
	for path, pathItem := range v.Swagger.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			if handler, ok := handlers[operation.OperationID]; ok {
				log.Printf("Registering route: %s %s (operationId: %s)", method, path, operation.OperationID)
				r.Handle(method, path, handler)
			}
		}
	}

	// Wrap Gin engine with validator middleware.
	handler := v.Middleware(r)

	log.Println("Server starting on :8081")
	log.Println("Swagger UI available at http://localhost:8081/docs")
	if err := http.ListenAndServe(":8081", handler); err != nil {
		log.Fatal(err)
	}
}
