package main

import (
	"sync"
)

// Storage represents the in-memory key-value storage for URL mappings
type Storage struct {
	shortToLong map[string]string // short_code -> long_url
	longToShort  map[string]string // long_url -> short_code
	mutex        sync.RWMutex
}

// NewStorage creates a new storage instance
func NewStorage() *Storage {
	return &Storage{
		shortToLong: make(map[string]string),
		longToShort:  make(map[string]string),
	}
}

// Store stores a mapping between short code and long URL
func (s *Storage) Store(shortCode, longURL string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	s.shortToLong[shortCode] = longURL
	s.longToShort[longURL] = shortCode
}

// GetByShort retrieves the long URL for a given short code
func (s *Storage) GetByShort(shortCode string) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	longURL, exists := s.shortToLong[shortCode]
	return longURL, exists
}

// GetByLong retrieves the short code for a given long URL
func (s *Storage) GetByLong(longURL string) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	shortCode, exists := s.longToShort[longURL]
	return shortCode, exists
}

// Exists checks if a short code exists
func (s *Storage) Exists(shortCode string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	_, exists := s.shortToLong[shortCode]
	return exists
} 
