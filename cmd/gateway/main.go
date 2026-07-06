package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// HealthResponse represents the JSON structure for our health check
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	// 1. Initialize the Multiplexer (Router)
	mux := http.NewServeMux()

	// 2. Register Endpoints
	mux.HandleFunc("/health", healthCheckHandler)

	// 3. Configure and Start the Server
	port := ":8080"
	log.Printf("Starting Resilient API Gateway on port %s...\n", port)
	
	// http.ListenAndServe blocks and keeps the server running
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}

// healthCheckHandler proves the server is alive (Crucial for Docker & AWS later)
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Build the response
	response := HealthResponse{
		Status:  "success",
		Message: "Gateway is up and running.",
	}

	// Safely encode and send the JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding health check response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}