package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Vaishnavi9843/engineering-command-center/internal/services"
)

func DashboardSummaryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	summary := services.GetDashboardSummary()

	json.NewEncoder(w).Encode(summary)
}