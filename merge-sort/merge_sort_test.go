package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "already sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "random array",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
		{
			name:     "array with duplicates",
			input:    []int{3, 3, 3, 1, 1, 2, 2},
			expected: []int{1, 1, 2, 2, 3, 3, 3},
		},
		{
			name:     "negative numbers",
			input:    []int{-3, 1, -4, 1, -5, 9, -2, 6},
			expected: []int{-5, -4, -3, -2, 1, 1, 6, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeSort(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMergeSortStability(t *testing.T) {
	// Test that merge sort is stable (preserves relative order of equal elements)
	type item struct {
		value int
		index int
	}

	input := []item{
		{value: 3, index: 0},
		{value: 1, index: 1},
		{value: 3, index: 2},
		{value: 1, index: 3},
	}

	// Custom comparison function for stability test
	less := func(a, b item) bool {
		return a.value < b.value
	}

	result := MergeSortStable(input, less)

	// Check that elements with same value maintain their relative order
	if result[0].index != 1 || result[1].index != 3 || result[2].index != 0 || result[3].index != 2 {
		t.Errorf("MergeSortStable failed to maintain stability: got %v", result)
	}
} 
