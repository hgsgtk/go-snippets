package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Handler handles HTTP requests for the URL shortener
type Handler struct {
	service *ShortenerService
}

// NewHandler creates a new handler instance
func NewHandler(service *ShortenerService) *Handler {
	return &Handler{
		service: service,
	}
}

// ShortenHandler handles POST /shorten requests
func (h *Handler) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Parse request body
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	// Validate request
	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}
	
	// Validate URL length (max 2048 characters)
	if len(req.URL) > 2048 {
		http.Error(w, "URL length must be less than 2048 characters", http.StatusBadRequest)
		return
	}
	
	// Create short URL
	shortURL, err := h.service.Shorten(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// Prepare response
	response := ShortenResponse{
		ShortURL: shortURL,
	}
	
	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// RedirectHandler handles GET /{short_code} requests
func (h *Handler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Extract short code from URL path
	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "" {
		http.Error(w, "Short code is required", http.StatusBadRequest)
		return
	}
	
	// Construct short URL for lookup
	shortURL := h.service.GetBaseURL() + path
	
	// Get original URL
	longURL, err := h.service.GetOriginal(shortURL)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}
	
	// Redirect to original URL
	http.Redirect(w, r, longURL, http.StatusFound)
} 
