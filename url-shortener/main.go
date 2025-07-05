package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

func main() {
	// Parse CLI arguments
	baseURL := flag.String("base-url", "https://short.url/", "Base URL for short URLs")
	port := flag.String("port", "8080", "Port to listen on")
	cleanupInterval := flag.Duration("cleanup-interval", 1*time.Hour, "Interval for cleaning up expired URLs")
	flag.Parse()
	
	// Create service and handler
	service := NewShortenerService(*baseURL)
	handler := NewHandler(service)
	
	// Start cleanup process for memory storage
	if memoryStorage, ok := service.storage.(*MemoryStorage); ok {
		memoryStorage.StartCleanup(*cleanupInterval)
		log.Printf("Started cleanup process with interval: %v", *cleanupInterval)
	}
	
	// Set up routing
	http.HandleFunc("/shorten", handler.ShortenHandler)
	http.HandleFunc("/", handler.RedirectHandler)
	
	// Start server
	log.Printf("URL Shortener server starting on :%s with base URL: %s", *port, *baseURL)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
} 
