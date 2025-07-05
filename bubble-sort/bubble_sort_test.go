package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Test case 1: [5,2,3,1]",
			input:    []int{5, 2, 3, 1},
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "Test case 2: [5,1,1,2,0,0]",
			input:    []int{5, 1, 1, 2, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the input to avoid modifying the original
			input := make([]int, len(tt.input))
			copy(input, tt.input)
			
			result := BubbleSort(input)
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSort(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
} 
