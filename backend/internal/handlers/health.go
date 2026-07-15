package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Vaishnavi9843/engineering-command-center/internal/services"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	response := services.Health()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}