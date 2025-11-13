package main

import (
	"go-demo-app/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", handlers.HealthHandler)

	log.Println("Starting server on :8081...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
