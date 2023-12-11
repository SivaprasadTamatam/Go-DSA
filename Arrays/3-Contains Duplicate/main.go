package main

import (
	"fmt"
	"sort"
)

// Solution 1: Brute force
func brute_force_containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// Solution 2: Sorting and Checking Neighboring Elements
func sort_containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

// Solution 3: Using map
func map_containsDuplicate(nums []int) bool {
	mp := make(map[int]bool)

	for _, num := range nums {
		if mp[num] {
			return true
		}
		mp[num] = true
	}

	return false
}

func main() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 3, 4}
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	fmt.Println("Solution 1 -  Brute Force:")
	fmt.Println("Solution 1 -  Brute Force:", brute_force_containsDuplicate(nums1)) // Output: true
	fmt.Println("Solution 1 -  Brute Force:", brute_force_containsDuplicate(nums2)) // Output: false
	fmt.Println("Solution 3 -  Brute Force:", brute_force_containsDuplicate(nums3)) // Output: true

	fmt.Println()

	fmt.Println("Solution 2 - Sorting:")
	fmt.Println("Solution 2 - Sorting:", sort_containsDuplicate(nums1)) // Output: true
	fmt.Println("Solution 2 - Sorting:", sort_containsDuplicate(nums2)) // Output: false
	fmt.Println("Solution 2 - Sorting:", sort_containsDuplicate(nums3)) // Output: true

	fmt.Println()

	fmt.Println("Solution 3 - Map:")
	fmt.Println("Solution 3 - Map:", map_containsDuplicate(nums1)) // Output: true
	fmt.Println("Solution 3 - Map:", map_containsDuplicate(nums2)) // Output: false
	fmt.Println("Solution 3 - Map:", map_containsDuplicate(nums3)) // Output: true
}
