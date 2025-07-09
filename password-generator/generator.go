package main

import (
	"math/rand"
	"time"
)

const (
	chars   = "abcdefghijklmnopqrstuvwxyz"
	symbols = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

// GeneratePassword generates a password with specified length and optional symbols
func GeneratePassword(length int, includeSymbols bool) string {
	// Create a new random source with current time
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	// Generate random password
	result := make([]byte, length)
	
	if includeSymbols {
		// Ensure at least one symbol is included
		symbolIndex := r.Intn(length)
		result[symbolIndex] = symbols[r.Intn(len(symbols))]
		
		// Fill the rest with random characters from the combined set
		charSet := chars + symbols
		for i := 0; i < length; i++ {
			if i != symbolIndex {
				result[i] = charSet[r.Intn(len(charSet))]
			}
		}
	} else {
		// Generate password with only lowercase letters
		for i := 0; i < length; i++ {
			result[i] = chars[r.Intn(len(chars))]
		}
	}
	
	return string(result)
} 
