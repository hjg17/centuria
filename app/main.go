package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)
// Demo API key now removed

func main() {
	// Health check endpoint
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// API endpoint that requires an API key for access
	http.HandleFunc("/api/v1/data", func(w http.ResponseWriter, r *http.Request) {
		expected := os.Getenv("API_SECRET")

		// Check for the API key in the Authorization header
		token := r.Header.Get("Authorization")

		// Validate the token against the expected value
		if expected == "" || token != expected {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Access granted")
	})

	http.ListenAndServe(":8080", nil)
}
