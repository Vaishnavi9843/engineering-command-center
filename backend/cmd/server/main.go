package main

import (
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`{
		"status":"healthy"
	}`))
}

func main() {

	http.HandleFunc("/api/health", healthHandler)

	log.Println("Server running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}