package handlers

import (
	"fmt"
	"net/http"
)

// HandlerFunc responds to HTTP requests.
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!!!")
}