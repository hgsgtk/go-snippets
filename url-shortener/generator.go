package main

import (
	"sync"
)

const (
	base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	codeLength  = 6
)

// Generator creates unique short codes using base62 encoding
type Generator struct {
	counter int64
	mutex   sync.Mutex
}

// NewGenerator creates a new generator instance
func NewGenerator() *Generator {
	return &Generator{
		counter: 0,
	}
}

// Generate creates a new unique short code
func (g *Generator) Generate() string {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	
	g.counter++
	return g.encodeBase62(g.counter)
}

// encodeBase62 converts a number to base62 string
func (g *Generator) encodeBase62(num int64) string {
	if num == 0 {
		return string(base62Chars[0])
	}
	
	var result string
	base := int64(len(base62Chars))
	
	for num > 0 {
		result = string(base62Chars[num%base]) + result
		num /= base
	}
	
	// Pad with leading zeros to ensure 6 characters
	for len(result) < codeLength {
		result = string(base62Chars[0]) + result
	}
	
	return result
} 
