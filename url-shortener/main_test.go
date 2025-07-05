package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

// Test data models are now defined in model.go

// Test storage functionality
func TestStorage(t *testing.T) {
	t.Parallel()
	storage := NewStorage()

	// Test storing and retrieving
	longURL := "https://www.google.com"
	shortCode := "abc123"
	
	storage.Store(shortCode, longURL)
	
	// Test GetByShort
	retrieved, exists := storage.GetByShort(shortCode)
	if !exists {
		t.Fatal("Expected to find short code")
	}
	if retrieved != longURL {
		t.Fatalf("Expected %s, got %s", longURL, retrieved)
	}

	// Test GetByLong
	retrievedCode, exists := storage.GetByLong(longURL)
	if !exists {
		t.Fatal("Expected to find long URL")
	}
	if retrievedCode != shortCode {
		t.Fatalf("Expected %s, got %s", shortCode, retrievedCode)
	}

	// Test non-existent short code
	_, exists = storage.GetByShort("nonexistent")
	if exists {
		t.Fatal("Expected not to find non-existent short code")
	}

	// Test non-existent long URL
	_, exists = storage.GetByLong("https://nonexistent.com")
	if exists {
		t.Fatal("Expected not to find non-existent long URL")
	}
}

// Test concurrent access to storage
func TestStorageConcurrency(t *testing.T) {
	t.Parallel()
	storage := NewStorage()
	var wg sync.WaitGroup
	
	// Test concurrent writes
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			shortCode := fmt.Sprintf("code%d", i)
			longURL := fmt.Sprintf("https://example%d.com", i)
			storage.Store(shortCode, longURL)
		}(i)
	}
	
	wg.Wait()
	
	// Verify all entries were stored
	for i := 0; i < 100; i++ {
		shortCode := fmt.Sprintf("code%d", i)
		longURL := fmt.Sprintf("https://example%d.com", i)
		
		retrieved, exists := storage.GetByShort(shortCode)
		if !exists {
			t.Fatalf("Expected to find short code %s", shortCode)
		}
		if retrieved != longURL {
			t.Fatalf("Expected %s, got %s", longURL, retrieved)
		}
	}
}

// Test short code generator
func TestGenerator(t *testing.T) {
	t.Parallel()
	generator := NewGenerator()
	
	// Test multiple generations
	codes := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		code := generator.Generate()
		
		// Check length
		if len(code) != 6 {
			t.Fatalf("Expected 6 characters, got %d", len(code))
		}
		
		// Check uniqueness
		if codes[code] {
			t.Fatalf("Duplicate code generated: %s", code)
		}
		codes[code] = true
		
		// Check if it's base62 (alphanumeric)
		for _, char := range code {
			if !((char >= '0' && char <= '9') || 
				 (char >= 'a' && char <= 'z') || 
				 (char >= 'A' && char <= 'Z')) {
				t.Fatalf("Invalid character in code %s: %c", code, char)
			}
		}
	}
}

// Test URL shortener service
func TestShortenerService(t *testing.T) {
	t.Parallel()
	service := NewShortenerService("https://short.url/")
	
	// Test creating new short URL
	longURL := "https://www.google.com"
	shortURL, err := service.Shorten(longURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if shortURL == "" {
		t.Fatal("Expected non-empty short URL")
	}
	
	// Test duplicate URL returns same short code
	shortURL2, err := service.Shorten(longURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if shortURL != shortURL2 {
		t.Fatalf("Expected same short URL for duplicate long URL")
	}
	
	// Test retrieving original URL
	retrieved, err := service.GetOriginal(shortURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if retrieved != longURL {
		t.Fatalf("Expected %s, got %s", longURL, retrieved)
	}
	
	// Test non-existent short URL
	_, err = service.GetOriginal("nonexistent")
	if err == nil {
		t.Fatal("Expected error for non-existent short URL")
	}
}

// Test HTTP handlers
func TestShortenHandler(t *testing.T) {
	t.Parallel()
	service := NewShortenerService("https://short.url/")
	handler := NewHandler(service)

	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		validateResponse func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "valid request",
			requestBody: ShortenRequest{URL: "https://www.google.com"},
			expectedStatus: http.StatusCreated,
			validateResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response ShortenResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				if err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}
				if response.ShortURL == "" {
					t.Fatal("Expected non-empty short URL in response")
				}
			},
		},
		{
			name: "invalid JSON",
			requestBody: "invalid json",
			expectedStatus: http.StatusBadRequest,
			validateResponse: func(t *testing.T, w *httptest.ResponseRecorder) {},
		},
		{
			name: "missing URL",
			requestBody: ShortenRequest{URL: ""},
			expectedStatus: http.StatusBadRequest,
			validateResponse: func(t *testing.T, w *httptest.ResponseRecorder) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			var err error
			
			switch v := tt.requestBody.(type) {
			case string:
				body = []byte(v)
			default:
				body, err = json.Marshal(tt.requestBody)
				if err != nil {
					t.Fatalf("Failed to marshal request body: %v", err)
				}
			}

			req := httptest.NewRequest("POST", "/shorten", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			handler.ShortenHandler(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			tt.validateResponse(t, w)
		})
	}
}

func TestRedirectHandler(t *testing.T) {
	t.Parallel()
	service := NewShortenerService("https://short.url/")
	handler := NewHandler(service)

	// Create a short URL for testing
	longURL := "https://www.google.com"
	shortURL, _ := service.Shorten(longURL)
	shortCode := strings.TrimPrefix(shortURL, "https://short.url/")

	tests := []struct {
		name           string
		path          string
		expectedStatus int
		expectedURL    string
	}{
		{
			name:           "successful redirect",
			path:          "/" + shortCode,
			expectedStatus: http.StatusFound,
			expectedURL:    longURL,
		},
		{
			name:           "non-existent short code",
			path:          "/nonexistent",
			expectedStatus: http.StatusNotFound,
			expectedURL:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			w := httptest.NewRecorder()

			handler.RedirectHandler(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedURL != "" {
				location := w.Header().Get("Location")
				if location != tt.expectedURL {
					t.Errorf("Expected Location header %s, got %s", tt.expectedURL, location)
				}
			}
		})
	}
}

// Test end-to-end functionality
func TestEndToEnd(t *testing.T) {
	t.Parallel()
	service := NewShortenerService("https://short.url/")
	handler := NewHandler(service)
	
	// Create short URL via handler
	reqBody := ShortenRequest{URL: "https://www.example.com"}
	jsonBody, _ := json.Marshal(reqBody)
	
	req := httptest.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	
	handler.ShortenHandler(w, req)
	
	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201, got %d", w.Code)
	}
	
	var response ShortenResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	
	// Extract short code and test redirect
	shortCode := strings.TrimPrefix(response.ShortURL, "https://short.url/")
	
	req = httptest.NewRequest("GET", "/"+shortCode, nil)
	w = httptest.NewRecorder()
	
	handler.RedirectHandler(w, req)
	
	if w.Code != http.StatusFound {
		t.Fatalf("Expected status 302, got %d", w.Code)
	}
	
	location := w.Header().Get("Location")
	if location != "https://www.example.com" {
		t.Fatalf("Expected Location header https://www.example.com, got %s", location)
	}
} 
