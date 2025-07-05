package main

// ShortenRequest represents the request body for the /shorten endpoint
type ShortenRequest struct {
	URL string `json:"url"`
}

// ShortenResponse represents the response body for the /shorten endpoint
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
} 
