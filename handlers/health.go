package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	res := HealthResponse{Status: "OK"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
