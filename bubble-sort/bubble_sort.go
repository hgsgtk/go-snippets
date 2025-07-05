package main

// BubbleSort sorts a slice of integers using the bubble sort algorithm
// and returns a new sorted slice without modifying the original
func BubbleSort(nums []int) []int {
	// Create a copy of the input slice
	result := make([]int, len(nums))
	copy(result, nums)
	
	n := len(result)
	
	// Outer loop: number of passes needed
	for i := 0; i < n-1; i++ {
		// Inner loop: compare adjacent elements
		for j := 0; j < n-1-i; j++ {
			// If current element is greater than next element, swap them
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	
	return result
} 



