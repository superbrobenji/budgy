// Depricated: This file is no longer used. It was used to start the server before but moved to a microservice architecture
package main

import (
	"net/http"

	transport "github.com/superbrobenji/budgy/infrastructure/transport/http"
)

func main() {
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := transport.NewServer()
	// Run the server
	http.ListenAndServe(":8080", mux)

}
