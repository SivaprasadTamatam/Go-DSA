package main

import "fmt"

func findMinBruteForce(nums []int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		// Check if the right half is sorted
		if nums[mid] < nums[r] {
			r = mid // Minimum is in the left half or at mid
		} else {
			l = mid + 1 // Minimum is in the right half
		}
	}
	return nums[l]
}

func main() {
	nums1 := []int{3, 4, 5, 1, 2}
	result1 := findMinBruteForce(nums1)
	fmt.Println("Minimum element (Brute Force):", result1)

	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	result2 := findMin(nums2)
	fmt.Println("Minimum element (Binary Search):", result2)
}
