package main

import (
	"fmt"
	"log"
	"net/http"
	"lamlados/backend/internal/router"
)


func main() {
	// Initialize the router.
	r := router.NewRouter()
	
	// Start the HTTP server on port 8080 and use the router.
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}