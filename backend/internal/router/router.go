package router

import (
	"net/http"

	"github.com/Vaishnavi9843/engineering-command-center/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/api/health", handlers.HealthHandler)
	mux.HandleFunc("/api/dashboard/summary", handlers.DashboardSummaryHandler)

}