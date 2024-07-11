package router

import (
	"lamlados/backend/internal/handler"
	"github.com/gorilla/mux" // Assuming you're using gorilla/mux for routing
)

// NewRouter initializes and returns a new router.
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Register the handler function for the root path.
	r.HandleFunc("/", handlers.HandlerFunc).Methods("GET")

	// You can add more routes here.

	return r
}