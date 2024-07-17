package main

import (
	"fmt"
	"sort"
)

// Helper function to sort a string
func sortString(s string) string {
	// Convert string to a slice of runes for sorting
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	// Convert sorted slice of runes back to a string
	return string(runes)
}

// Function to check if two strings are anagrams using brute force approach
func isAnagramBruteForce(s string, t string) bool {
	// If lengths are not equal, they cannot be anagrams
	if len(s) != len(t) {
		return false
	}
	// Sort both strings and compare
	return sortString(s) == sortString(t)
}

func isAnagram(s string, t string) bool {

	if len(t) != len(s) {
		return false
	}

	mp := make(map[rune]int)

	for _, ch := range s {
		mp[ch]++
	}

	for _, ch := range t {
		mp[ch]--

		if mp[ch] < 0 {
			return false
		}
	}

	for _, c := range mp {
		if c != 0 {
			return false
		}

	}

	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	result := isAnagramBruteForce(s, t)
	if result {
		fmt.Println("The strings are anagrams.")
	} else {
		fmt.Println("The strings are not anagrams.")
	}

	result = isAnagram(s, t)
	if result {
		fmt.Println("The strings are anagrams.")
	} else {
		fmt.Println("The strings are not anagrams.")
	}
}
