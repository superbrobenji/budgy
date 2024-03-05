package transport

import (
	"fmt"
	"net/http"

	"github.com/superbrobenji/budgy/core/aggregate"
	categoryDatastore "github.com/superbrobenji/budgy/infrastructure/persistence/dataStore/category"
)

func NewServer() (mux *http.ServeMux) {
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	muxServer := http.NewServeMux()

	// Register the routes and handlers
	muxServer.Handle("/", &homeHandler{})
	muxServer.Handle("/api/v1.0", &categoryHandler{})

	return muxServer
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

type categoryHandler struct{}

func (h *categoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, _ := aggregate.NewCategory("Groceries")
	database := categoryDatastore.NewDynamoCategoryRepository()
    err := database.CreateCategory(&category)
    if err != nil {
        fmt.Println(err)
        w.Write([]byte("Error creating category"))
        return
    }
    w.Write([]byte("This is the category page"))
}
