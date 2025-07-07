package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Example 1: Target exists",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			name:     "Example 2: Target does not exist",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Search(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("Search(%v, %d) = %d, want %d", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
} 
