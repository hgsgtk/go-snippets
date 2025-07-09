package main

import (
	"strings"
	"testing"
	"unicode"
)


func TestGeneratePassword(t *testing.T) {
	t.Run("should generate password with specified length", func(t *testing.T) {
		testCases := []int{8, 16, 32, 50, 100}
		for _, length := range testCases {
			password := GeneratePassword(length, false)
			if len(password) != length {
				t.Errorf("Expected password length %d, got %d", length, len(password))
			}
		}
	})

	t.Run("should generate only lowercase alphabets when symbols disabled", func(t *testing.T) {
		password := GeneratePassword(20, false)
		for _, char := range password {
			if !unicode.IsLower(char) || !unicode.IsLetter(char) {
				t.Errorf("Expected only lowercase letters, got %c", char)
			}
		}
	})

	t.Run("should generate different passwords on multiple calls", func(t *testing.T) {
		passwords := make(map[string]bool)
		for i := 0; i < 100; i++ {
			password := GeneratePassword(15, false)
			if passwords[password] {
				t.Errorf("Duplicate password generated: %s", password)
			}
			passwords[password] = true
		}
	})

	t.Run("should contain only a-z characters when symbols disabled", func(t *testing.T) {
		password := GeneratePassword(25, false)
		validChars := "abcdefghijklmnopqrstuvwxyz"
		for _, char := range password {
			if !strings.ContainsRune(validChars, char) {
				t.Errorf("Invalid character in password: %c", char)
			}
		}
	})

	t.Run("should include symbols when symbols enabled", func(t *testing.T) {
		password := GeneratePassword(20, true)
		hasSymbol := false
		for _, char := range password {
			if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
				hasSymbol = true
				break
			}
		}
		if !hasSymbol {
			t.Errorf("Password should contain at least one symbol when symbols enabled")
		}
	})

	t.Run("should guarantee at least one symbol when symbols enabled", func(t *testing.T) {
		// Test multiple times to ensure the algorithm guarantees at least one symbol
		for i := 0; i < 100000; i++ {
			password := GeneratePassword(8, true)
			hasSymbol := false
			for _, char := range password {
				if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
					hasSymbol = true
					break
				}
			}
			if !hasSymbol {
				t.Errorf("Password generated without symbols: %s", password)
			}
		}
	})

	t.Run("should contain at least one symbol in short passwords", func(t *testing.T) {
		// Test with minimum length to ensure symbols are included even in short passwords
		password := GeneratePassword(8, true)
		hasSymbol := false
		for _, char := range password {
			if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
				hasSymbol = true
				break
			}
		}
		if !hasSymbol {
			t.Errorf("Short password should contain at least one symbol: %s", password)
		}
	})

	t.Run("should contain valid characters when symbols enabled", func(t *testing.T) {
		password := GeneratePassword(25, true)
		validChars := "abcdefghijklmnopqrstuvwxyz!@#$%^&*()_+-=[]{}|;:,.<>?"
		for _, char := range password {
			if !strings.ContainsRune(validChars, char) {
				t.Errorf("Invalid character in password: %c", char)
			}
		}
	})

	t.Run("should generate different passwords with symbols on multiple calls", func(t *testing.T) {
		passwords := make(map[string]bool)
		for i := 0; i < 100; i++ {
			password := GeneratePassword(15, true)
			if passwords[password] {
				t.Errorf("Duplicate password generated: %s", password)
			}
			passwords[password] = true
		}
	})
} 
