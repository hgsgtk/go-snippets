package main

import (
	"errors"
	"net/url"
	"strings"
)

// ShortenerService handles URL shortening business logic
type ShortenerService struct {
	storage   *Storage
	generator *Generator
	baseURL   string
}

// NewShortenerService creates a new shortener service
func NewShortenerService(baseURL string) *ShortenerService {
	return &ShortenerService{
		storage:   NewStorage(),
		generator: NewGenerator(),
		baseURL:   baseURL,
	}
}

// Shorten creates a short URL from a long URL
func (s *ShortenerService) Shorten(longURL string) (string, error) {
	// Validate URL
	if err := s.validateURL(longURL); err != nil {
		return "", err
	}
	
	// Check if URL already exists
	if existingCode, exists := s.storage.GetByLong(longURL); exists {
		return s.baseURL + existingCode, nil
	}
	
	// Generate new short code
	shortCode := s.generator.Generate()
	
	// Store the mapping
	s.storage.Store(shortCode, longURL)
	
	return s.baseURL + shortCode, nil
}

// GetOriginal retrieves the original URL from a short URL
func (s *ShortenerService) GetOriginal(shortURL string) (string, error) {
	// Extract short code from URL
	shortCode := s.extractShortCode(shortURL)
	if shortCode == "" {
		return "", errors.New("invalid short URL format")
	}
	
	// Look up the original URL
	longURL, exists := s.storage.GetByShort(shortCode)
	if !exists {
		return "", errors.New("short URL not found")
	}
	
	return longURL, nil
}

// GetBaseURL returns the base URL used by the service
func (s *ShortenerService) GetBaseURL() string {
	return s.baseURL
}

// validateURL checks if the provided URL is valid
func (s *ShortenerService) validateURL(urlStr string) error {
	if urlStr == "" {
		return errors.New("URL cannot be empty")
	}
	
	// Add scheme if missing
	if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
		urlStr = "https://" + urlStr
	}
	
	_, err := url.Parse(urlStr)
	if err != nil {
		return errors.New("invalid URL format")
	}
	
	return nil
}

// extractShortCode extracts the short code from a short URL
func (s *ShortenerService) extractShortCode(shortURL string) string {
	if !strings.HasPrefix(shortURL, s.baseURL) {
		return ""
	}
	
	shortCode := strings.TrimPrefix(shortURL, s.baseURL)
	if len(shortCode) != 6 {
		return ""
	}
	
	return shortCode
} 
