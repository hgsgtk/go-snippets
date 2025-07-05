package main

import (
	"crypto/rand"
	"sync"
)

const (
	base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	codeLength  = 6
)

// Generator creates unique short codes using random base62 strings
type Generator struct {
	usedCodes map[string]bool
	mutex     sync.RWMutex
}

// NewGenerator creates a new generator instance
func NewGenerator() *Generator {
	return &Generator{
		usedCodes: make(map[string]bool),
	}
}

// Generate creates a new unique short code
func (g *Generator) Generate() string {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	
	// Generate random codes until we find an unused one
	for {
		code := g.generateRandomCode()
		if !g.usedCodes[code] {
			g.usedCodes[code] = true
			return code
		}
	}
}

// generateRandomCode creates a random 6-character base62 string
func (g *Generator) generateRandomCode() string {
	// Generate random bytes
	bytes := make([]byte, codeLength)
	rand.Read(bytes)
	
	// Convert to base62 string
	var result string
	for _, b := range bytes {
		index := int(b) % len(base62Chars)
		result += string(base62Chars[index])
	}
	
	return result
} 
