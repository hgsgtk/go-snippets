package main

import (
	"fmt"
	"sort"
)

// bubbleSort implements the bubble sort algorithm
func bubbleSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)
	
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// quickSort implements the quick sort algorithm
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	pivot := arr[0]
	var left, right []int
	
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	
	left = quickSort(left)
	right = quickSort(right)
	
	return append(append(left, pivot), right...)
}

// mergeSort implements the merge sort algorithm
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	
	return merge(left, right)
}

// merge helper function for merge sort
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	
	return result
}

func main() {
	// Sample arrays to sort
	arrays := [][]int{
		{64, 34, 25, 12, 22, 11, 90},
		{5, 2, 4, 6, 1, 3},
		{1},
		{},
		{3, 3, 3, 3},
	}
	
	for i, arr := range arrays {
		fmt.Printf("Array %d: %v\n", i+1, arr)
		
		if len(arr) == 0 {
			fmt.Println("  Empty array, nothing to sort")
			fmt.Println()
			continue
		}
		
		// Sort using different methods
		bubbleResult := bubbleSort(arr)
		quickResult := quickSort(arr)
		mergeResult := mergeSort(arr)
		
		// Built-in sort for comparison
		builtinResult := make([]int, len(arr))
		copy(builtinResult, arr)
		sort.Ints(builtinResult)
		
		fmt.Printf("  Bubble Sort: %v\n", bubbleResult)
		fmt.Printf("  Quick Sort:  %v\n", quickResult)
		fmt.Printf("  Merge Sort:  %v\n", mergeResult)
		fmt.Printf("  Built-in:    %v\n", builtinResult)
		fmt.Println()
	}
} 
