package main

import (
	"fmt"
	"math"
)

// Brute force approach
func maxSubArrayBruteForce(nums []int) int {
	ms := math.MinInt32
	for i := 0; i < len(nums); i++ {
		cs := 0
		for j := i; j < len(nums); j++ {
			cs += nums[j]
			ms = max(ms, cs)
		}
	}
	return ms
}

// Kadane's algorithm
func maxSubArrayKadane(nums []int) int {
	ms := math.MinInt32
	cs := 0

	for i := 0; i < len(nums); i++ {
		cs = max(nums[i], cs+nums[i])
		ms = max(cs, ms)
	}

	return ms
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums2 := []int{1}
	nums3 := []int{5, 4, -1, 7, 8}

	fmt.Println("Using Brute Force approach:")
	fmt.Println(maxSubArrayBruteForce(nums1)) // Output: 6
	fmt.Println(maxSubArrayBruteForce(nums2)) // Output: 1
	fmt.Println(maxSubArrayBruteForce(nums3)) // Output: 23

	fmt.Println("\nUsing Kadane's Algorithm:")
	fmt.Println(maxSubArrayKadane(nums1)) // Output: 6
	fmt.Println(maxSubArrayKadane(nums2)) // Output: 1
	fmt.Println(maxSubArrayKadane(nums3)) // Output: 23
}
