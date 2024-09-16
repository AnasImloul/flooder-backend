package main

import (
	"flood-backend/src"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// CORS middleware function to allow all origins
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow specific HTTP methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// Allow specific headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		// Pass to the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Create a new main router
	r := mux.NewRouter()

	// Create the sub-router
	floodRouter := r.PathPrefix("/flood").Subrouter()

	// Register flood routes with the floodRouter
	src.RegisterRoutes(floodRouter)

	// Apply CORS middleware to the main router
	handlerWithCORS := enableCORS(r)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", handlerWithCORS))
}
