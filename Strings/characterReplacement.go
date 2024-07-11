package main

import (
	"fmt"
)

// Function to find the length of the longest substring with the same character
// after performing at most k replacements
func characterReplacementBF(s string, k int) int {
	n := len(s)
	maxLength := 0

	// Iterate over all possible starting points of substrings
	for start := 0; start < n; start++ {
		// Iterate over all possible ending points
		for end := start; end < n; end++ {
			fmt.Println(string(s[start : end+1]))
			// Frequency map to count the frequency of each character in the current substring
			freq := make(map[byte]int)
			for i := start; i <= end; i++ {
				freq[s[i]]++
			}

			// Find the most frequent character in the current substring
			maxFreq := 0
			for _, count := range freq {
				if count > maxFreq {
					maxFreq = count
				}
			}

			// Calculate the number of replacements needed
			// to make all characters in the substring the same
			replacements := (end - start + 1) - maxFreq

			fmt.Println(maxFreq, replacements)
			// If replacements are within the allowed limit, update maxLength
			if replacements <= k {
				if (end - start + 1) > maxLength {
					maxLength = end - start + 1
					fmt.Println(maxLength, end, start, replacements)
				}
			}
		}
		fmt.Println()
	}

	return maxLength
}

// Function to find the length of the longest substring with the same character
// after performing at most k replacements
func characterReplacementOP(s string, k int) int {
	n := len(s)
	freq := make(map[byte]int)
	maxCount := 0
	maxLength := 0
	start := 0

	// Iterate over the string with the end pointer
	for end := 0; end < n; end++ {
		// Increment the count of the current character
		freq[s[end]]++
		// Update the maxCount to the highest frequency of a single character in the current window
		if freq[s[end]] > maxCount {
			maxCount = freq[s[end]]
		}

		// If the window size minus the count of the most frequent character is greater than k,
		// shrink the window from the left
		if (end-start+1)-maxCount > k {
			freq[s[start]]--
			start++

		}

		// Update the maximum length of the window
		maxLength = end - start + 1
	}

	return maxLength
}

func main() {
	s := "AABABBA"
	k := 2
	result := characterReplacementOP(s, k)
	fmt.Println("The length of the longest substring is:", result)
}
