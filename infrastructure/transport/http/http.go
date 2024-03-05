package transport

import (
	"net/http"
)

func NewServer() (mux *http.ServeMux) {
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	muxServer := http.NewServeMux()

	// Register the routes and handlers
	muxServer.Handle("/", &homeHandler{})

	return muxServer
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
