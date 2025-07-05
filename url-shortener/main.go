package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Parse CLI arguments
	baseURL := flag.String("base-url", "https://short.url/", "Base URL for short URLs")
	port := flag.String("port", "8080", "Port to listen on")
	flag.Parse()
	
	// Create service and handler
	service := NewShortenerService(*baseURL)
	handler := NewHandler(service)
	
	// Set up routing
	http.HandleFunc("/shorten", handler.ShortenHandler)
	http.HandleFunc("/", handler.RedirectHandler)
	
	// Start server
	log.Printf("URL Shortener server starting on :%s with base URL: %s", *port, *baseURL)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
} 
