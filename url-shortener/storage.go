package main

import (
	"sync"
	"time"
)

// StorageInterface defines the contract for URL storage implementations
type StorageInterface interface {
	Store(shortCode, longURL string, ttl *time.Duration) error
	GetByShort(shortCode string) (string, bool)
	GetByLong(longURL string) (string, bool)
	Exists(shortCode string) bool
	StartCleanup(interval time.Duration)
	StopCleanup()
}

// MemoryStorage implements in-memory key-value storage for URL mappings
type MemoryStorage struct {
	shortToLong map[string]*URLMapping // short_code -> URLMapping
	longToShort  map[string]string     // long_url -> short_code
	mutex        sync.RWMutex
	cleanupTicker *time.Ticker
	stopCleanup   chan bool
}

// NewMemoryStorage creates a new in-memory storage instance
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		shortToLong: make(map[string]*URLMapping),
		longToShort:  make(map[string]string),
		stopCleanup:  make(chan bool),
	}
}

// Store stores a mapping between short code and long URL with optional TTL
func (s *MemoryStorage) Store(shortCode, longURL string, ttl *time.Duration) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	now := time.Now()
	var expiresAt *time.Time
	
	if ttl != nil {
		expiry := now.Add(*ttl)
		expiresAt = &expiry
	}
	
	mapping := &URLMapping{
		ShortCode: shortCode,
		LongURL:   longURL,
		CreatedAt: now,
		ExpiresAt: expiresAt,
	}
	
	s.shortToLong[shortCode] = mapping
	s.longToShort[longURL] = shortCode
	return nil
}

// GetByShort retrieves the long URL for a given short code
func (s *MemoryStorage) GetByShort(shortCode string) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	mapping, exists := s.shortToLong[shortCode]
	if !exists {
		return "", false
	}
	
	// Check if expired
	if mapping.ExpiresAt != nil && time.Now().After(*mapping.ExpiresAt) {
		return "", false
	}
	
	return mapping.LongURL, true
}

// GetByLong retrieves the short code for a given long URL
func (s *MemoryStorage) GetByLong(longURL string) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	shortCode, exists := s.longToShort[longURL]
	if !exists {
		return "", false
	}
	
	// Check if expired
	mapping, exists := s.shortToLong[shortCode]
	if !exists {
		return "", false
	}
	
	if mapping.ExpiresAt != nil && time.Now().After(*mapping.ExpiresAt) {
		return "", false
	}
	
	return shortCode, true
}

// Exists checks if a short code exists and is not expired
func (s *MemoryStorage) Exists(shortCode string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	mapping, exists := s.shortToLong[shortCode]
	if !exists {
		return false
	}
	
	// Check if expired
	if mapping.ExpiresAt != nil && time.Now().After(*mapping.ExpiresAt) {
		return false
	}
	
	return true
}

// StartCleanup starts the background cleanup process
func (s *MemoryStorage) StartCleanup(interval time.Duration) {
	s.cleanupTicker = time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-s.cleanupTicker.C:
				s.cleanupExpired()
			case <-s.stopCleanup:
				return
			}
		}
	}()
}

// StopCleanup stops the background cleanup process
func (s *MemoryStorage) StopCleanup() {
	if s.cleanupTicker != nil {
		s.cleanupTicker.Stop()
	}
	close(s.stopCleanup)
}

// cleanupExpired removes expired entries from storage
func (s *MemoryStorage) cleanupExpired() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	now := time.Now()
	expiredCodes := []string{}
	
	// Find expired entries
	for shortCode, mapping := range s.shortToLong {
		if mapping.ExpiresAt != nil && now.After(*mapping.ExpiresAt) {
			expiredCodes = append(expiredCodes, shortCode)
		}
	}
	
	// Remove expired entries
	for _, shortCode := range expiredCodes {
		mapping := s.shortToLong[shortCode]
		delete(s.shortToLong, shortCode)
		delete(s.longToShort, mapping.LongURL)
	}
}

// StorageFactory creates storage instances based on configuration
type StorageFactory struct{}

// NewStorageFactory creates a new storage factory
func NewStorageFactory() *StorageFactory {
	return &StorageFactory{}
}

// CreateStorage creates a storage instance based on the storage type
func (f *StorageFactory) CreateStorage(storageType string) (StorageInterface, error) {
	switch storageType {
	case "memory":
		return NewMemoryStorage(), nil
	case "mysql":
		// TODO: Implement MySQL storage
		return nil, nil
	default:
		return NewMemoryStorage(), nil // Default to memory storage
	}
} 
