package main

import (
	"fmt"
	"net/http"
	"lamlados/backend/internal/handlers"
)


func main() {
	// Register the handler function for the root path.
	http.HandleFunc("/", handlers.HandlerFunc)
	
	// Start the HTTP server on port 8080.
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}