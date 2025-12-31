package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/getkin/kin-openapi/routers/legacy"
	validator "github.com/vihuvac/go-openapi-validator"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func main() {
	v, err := validator.New("openapi.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// For net/http, we need to explicitly set the legacy router.
	router, err := legacy.NewRouter(v.Swagger)
	if err != nil {
		log.Fatal(err)
	}
	validator.WithRouter(router)(v.Options)

	mux := http.NewServeMux()

	// Register Swagger UI.
	v.HandleSwaggerUI(mux)

	// API Handler.
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req HelloRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := HelloResponse{
			Message: "Hello " + req.Name,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	// Wrap with validator middleware.
	handler := v.Middleware(mux)

	log.Println("Server starting on :8081")
	log.Println("Swagger UI available at http://localhost:8081/docs")
	if err := http.ListenAndServe(":8081", handler); err != nil {
		log.Fatal(err)
	}
}
