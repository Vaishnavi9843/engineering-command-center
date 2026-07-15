package server

import (
	"net/http"

	"github.com/Vaishnavi9843/engineering-command-center/internal/router"
)

func New(address string) *http.Server {

	mux := http.NewServeMux()

	router.RegisterRoutes(mux)

	return &http.Server{
		Addr:    address,
		Handler: mux,
	}
}