package main

import (
	"fmt"
)

// ValidateLength validates that the password length is within the allowed range
func ValidateLength(length int) error {
	if length < 8 {
		return fmt.Errorf("length must be at least 8 characters")
	}
	if length > 100 {
		return fmt.Errorf("length must be at most 100 characters")
	}
	return nil
}
