package main

// MergeSort sorts an array of integers using the merge sort algorithm
// Time Complexity: O(n log n) - divides array in half each time, then merges
// Space Complexity: O(n) - needs extra space for temporary arrays
func MergeSort(arr []int) []int {
	// Base case 1: If array is empty, return empty array
	if len(arr) == 0 {
		return []int{}
	}
	// Base case 2: If array has only one element, it's already sorted
	// Return a copy to avoid modifying the original array
	if len(arr) == 1 {
		return append([]int(nil), arr...)
	}
	
	// Find the middle point to divide array into two halves
	mid := len(arr) / 2
	
	// Recursively sort the left half of the array
	// arr[:mid] creates a slice from index 0 to mid-1
	left := MergeSort(arr[:mid])
	
	// Recursively sort the right half of the array
	// arr[mid:] creates a slice from index mid to the end
	right := MergeSort(arr[mid:])
	
	// Merge the two sorted halves and return the result
	return merge(left, right)
}

// merge combines two sorted arrays into a single sorted array
// This is the key operation that makes merge sort work
func merge(left, right []int) []int {
	// Create a result slice with initial capacity to avoid reallocations
	// Capacity is sum of both input array lengths
	result := make([]int, 0, len(left)+len(right))
	
	// Initialize pointers for both arrays
	i, j := 0, 0
	
	// Compare elements from both arrays and merge them in sorted order
	// Continue until we've processed all elements from either array
	for i < len(left) && j < len(right) {		
		// If left element is smaller or equal, add it to result
		// Using <= ensures stability (equal elements maintain relative order)
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++ // Move to next element in left array
		} else {
			// If right element is smaller, add it to result
			result = append(result, right[j])
			j++ // Move to next element in right array
		}
	}
		
	// After the loop, one of the arrays may still have remaining elements
	// Add all remaining elements from left array (if any)
	// left[i:] creates a slice from index i to the end of left array
	result = append(result, left[i:]...)
	
	// Add all remaining elements from right array (if any)
	// right[j:] creates a slice from index j to the end of right array
	result = append(result, right[j:]...)
	
	return result
}

// MergeSortStable sorts an array using merge sort with a custom comparison function
// This version is stable, meaning it preserves the relative order of equal elements
func MergeSortStable[T any](arr []T, less func(T, T) bool) []T {
	if len(arr) == 0 {
		return []T{}
	}
	if len(arr) == 1 {
		return append([]T(nil), arr...)
	}
	mid := len(arr) / 2
	left := MergeSortStable(arr[:mid], less)
	right := MergeSortStable(arr[mid:], less)
	return mergeStable(left, right, less)
}

func mergeStable[T any](left, right []T, less func(T, T) bool) []T {
	result := make([]T, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if less(left[i], right[j]) {
			result = append(result, left[i])
			i++
		} else if less(right[j], left[i]) {
			result = append(result, right[j])
			j++
		} else {
			// Equal: take from left to preserve stability
			result = append(result, left[i])
			i++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
} 
