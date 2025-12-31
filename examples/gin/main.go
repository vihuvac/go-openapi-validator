package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/vihuvac/go-openapi-validator"
)

type ApiResponse struct {
	Message string `json:"message"`
}

func main() {
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

	// Regular (manual) route registration
	r.GET("/health/liveness", func(c *gin.Context) {
		c.JSON(http.StatusOK, ApiResponse{
			Message: "Ok",
		})
	})

	r.GET("/health/readiness", func(c *gin.Context) {
		c.JSON(http.StatusOK, ApiResponse{
			Message: "Ok",
		})
	})

	// Wrap Gin engine with validator middleware.
	handler := v.Middleware(r)

	log.Println("Server starting on :8081")
	log.Println("Swagger UI available at http://localhost:8081/docs")
	if err := http.ListenAndServe(":8081", handler); err != nil {
		log.Fatal(err)
	}
}
