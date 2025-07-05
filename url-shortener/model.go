package main

import "time"

// ShortenRequest represents the request body for the /shorten endpoint
type ShortenRequest struct {
	URL string `json:"url"`
	TTL *int64 `json:"ttl,omitempty"` // TTL in seconds (optional)
}

// ShortenResponse represents the response body for the /shorten endpoint
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"` // When the URL expires (if TTL was set)
}

// URLMapping represents a URL mapping with metadata
type URLMapping struct {
	ShortCode string
	LongURL   string
	CreatedAt time.Time
	ExpiresAt *time.Time
} 
