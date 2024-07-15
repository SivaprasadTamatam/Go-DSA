package main

import (
	"fmt"
	"math"
)

// Function to find the minimum window substring
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	// Create a frequency map for characters in t
	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	// Initialize variables for the sliding window
	start, end := 0, 0
	minStart, minLen := 0, math.MaxInt32
	needed := len(tFreq)
	have := 0
	windowFreq := make(map[byte]int)

	for end < len(s) {
		charEnd := s[end]
		windowFreq[charEnd]++

		if count, ok := tFreq[charEnd]; ok && windowFreq[charEnd] == count {
			have++
		}

		// Contract the window as much as possible while it still contains all characters in t
		for have == needed {
			if end-start+1 < minLen {
				minStart = start
				minLen = end - start + 1
			}

			charStart := s[start]
			windowFreq[charStart]--
			if count, ok := tFreq[charStart]; ok && windowFreq[charStart] < count {
				have--
			}
			start++
		}

		end++
	}

	if minLen == math.MaxInt32 {
		return ""
	}
	return s[minStart : minStart+minLen]
}

// Helper function to check if the substring contains all characters of t
func containsAll(sFreq, tFreq map[byte]int) bool {
	for key, val := range tFreq {
		if sFreq[key] < val {
			return false
		}
	}
	return true
}

// Function to find the minimum window substring using brute force
func minWindowBruteForce(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	minStart, minLen := 0, math.MaxInt32
	n := len(s)

	for start := 0; start < n; start++ {
		sFreq := make(map[byte]int)
		for end := start; end < n; end++ {
			sFreq[s[end]]++
			if containsAll(sFreq, tFreq) {
				if end-start+1 < minLen {
					minStart = start
					minLen = end - start + 1
				}
				break // No need to expand further as we've found a valid window
			}
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}
	return s[minStart : minStart+minLen]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	result := minWindow(s, t)
	fmt.Println("The minimum window substring is:", result)

	result = minWindowBruteForce(s, t)
	fmt.Println("The minimum window substring using brute force is:", result)
}
