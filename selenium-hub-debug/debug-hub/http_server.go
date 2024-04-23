package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

// RequestLogger is an HTTP handler that prints all the details of the request it receives
func RequestLogger(w http.ResponseWriter, r *http.Request) {
	// Print HTTP method, URL, and remote address
	fmt.Printf("Received %s request for %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)

  // Log the full query string
	if query := r.URL.Query(); len(query) > 0 {
		fmt.Printf("Query Parameters: %v\n", query)
	}

	// Optionally, you can print headers or other parts of the request
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}

  // Read and log the body
	// Make sure not to do this with large bodies or in production without limiting size
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}
	fmt.Printf("Body: %s\n", body)

	// It's important to close the body when you're done with it
	defer r.Body.Close()

  // Rewind the body if you plan to use it again after this handler (middleware scenario)
	r.Body = io.NopCloser(bytes.NewReader(body))


	// Set the header Content-Type for JSON response
  response := []byte(`{"success": true, "message": "Node registered successfully"}`)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func startServer(port string) {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Set up the route specific to this mux
	mux.HandleFunc("/", RequestLogger)

	// Start the HTTP server on a specific port with the new mux
	log.Printf("Starting server on :%s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Error starting server on port %s: %s\n", port, err)
	}
}
