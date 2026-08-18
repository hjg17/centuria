package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)
// Demo API key
const hardCodedKey = "AIzaEXAMPLED99NlL8x9A9T3tJ9h118Zc1gC1DK"

func main() {
	// Health check endpoint
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// API endpoint that requires an API key for access
	http.HandleFunc("/api/v1/data", func(w http.ResponseWriter, r *http.Request) {
		expected := os.Getenv("API_SECRET")
		if expected == "" {
			expected = hardCodedKey
		}

		// Check for the API key in the Authorization header
		token := r.Header.Get("Authorization")

		// Ideally, we would also check for empty tokens and deny authorization, but we will allow it to be empty for this demo
		if token != expected {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Access granted")
	})

	http.ListenAndServe(":8080", nil)
}
