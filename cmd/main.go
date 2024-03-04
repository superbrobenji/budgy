package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	transport "github.com/superbrobenji/budgy/infrastructure/transport/http"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := transport.NewServer()
	// Run the server
	http.ListenAndServe(":8080", mux)

}
