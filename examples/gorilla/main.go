package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	validator "github.com/vihuvac/go-openapi-validator"
)

type ApiResponse struct {
	Message string `json:"message"`
}

func main() {
	// New defaults to gorillamux router if no WithRouter option is provided.
	v, err := validator.New(
		"openapi.yaml",
		validator.WithValidateRequests(true),
		validator.WithValidateResponses(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	// Register Swagger UI using SwaggerUIHandler for Gorilla Mux.
	r.PathPrefix(v.Options.SwaggerUIPath).Handler(v.SwaggerUIHandler())

	// API Handler
	r.HandleFunc("/health/liveness", func(w http.ResponseWriter, r *http.Request) {
		resp := ApiResponse{
			Message: "Ok",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}).Methods("GET")

	r.HandleFunc("/health/readiness", func(w http.ResponseWriter, r *http.Request) {
		resp := ApiResponse{
			Message: "Ok",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}).Methods("GET")

	// Wrap with validator middleware.
	handler := v.Middleware(r)

	log.Println("Server starting on :8081")
	log.Println("Swagger UI available at http://localhost:8081/docs")
	if err := http.ListenAndServe(":8081", handler); err != nil {
		log.Fatal(err)
	}
}
