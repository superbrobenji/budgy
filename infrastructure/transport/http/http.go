package transport

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	services "github.com/superbrobenji/budgy/core/service/createItemTest"
)

func NewServer() (mux *http.ServeMux) {
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	muxServer := http.NewServeMux()

	// Register the routes and handlers
	muxServer.Handle("/api/v1.0", &itemHandler{})

	return muxServer
}

type itemHandler struct{}

func (h *itemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	itemService, err := services.NewItemService()
	if err != nil {
		w.Write([]byte("Error creating category"))
		return
	}
	item, error := itemService.CreateItem(uuid.New())
	if error != nil {
		w.Write([]byte("Error creating category"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}
