package main

import "fmt"

// Brute force approach
func searchBruteForce(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

// Optimized binary search approach
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	// Brute force solution
	resultBruteForce := searchBruteForce(nums, target)
	fmt.Println("Index of target using brute force:", resultBruteForce)

	// Optimized binary search solution
	resultOptimized := search(nums, target)
	fmt.Println("Index of target using binary search:", resultOptimized)
}
