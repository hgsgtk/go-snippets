package binarysearch

// Search performs a binary search on a sorted array of integers.
// Given an array of integers nums which is sorted in ascending order, and an integer target,
// this function searches for target in nums. If target exists, then return its index.
// Otherwise, return -1. The algorithm has O(log n) runtime complexity.
func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return -1
} 
