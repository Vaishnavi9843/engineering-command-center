package main

import (
	"log"

	"github.com/Vaishnavi9843/engineering-command-center/internal/server"
)

func main() {
	srv := server.New(":8080")

	log.Println("Server running on :8080")

	log.Fatal(srv.ListenAndServe())
}