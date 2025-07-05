package main

import (
	"sync"
)

// StorageInterface defines the contract for URL storage implementations
type StorageInterface interface {
	Store(shortCode, longURL string) error
	GetByShort(shortCode string) (string, bool)
	GetByLong(longURL string) (string, bool)
	Exists(shortCode string) bool
}

// MemoryStorage implements in-memory key-value storage for URL mappings
type MemoryStorage struct {
	shortToLong map[string]string // short_code -> long_url
	longToShort  map[string]string // long_url -> short_code
	mutex        sync.RWMutex
}

// NewMemoryStorage creates a new in-memory storage instance
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		shortToLong: make(map[string]string),
		longToShort:  make(map[string]string),
	}
}

// Store stores a mapping between short code and long URL
func (s *MemoryStorage) Store(shortCode, longURL string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	s.shortToLong[shortCode] = longURL
	s.longToShort[longURL] = shortCode
	return nil
}

// GetByShort retrieves the long URL for a given short code
func (s *MemoryStorage) GetByShort(shortCode string) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	longURL, exists := s.shortToLong[shortCode]
	return longURL, exists
}

// GetByLong retrieves the short code for a given long URL
func (s *MemoryStorage) GetByLong(longURL string) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	shortCode, exists := s.longToShort[longURL]
	return shortCode, exists
}

// Exists checks if a short code exists
func (s *MemoryStorage) Exists(shortCode string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	_, exists := s.shortToLong[shortCode]
	return exists
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
