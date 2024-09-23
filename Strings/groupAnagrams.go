package main

import (
	"fmt"
	"sort"
	"strings"
)

// O(klogk) here k is the length of word
func SortWord(s string) string {
	slice := strings.Split(s, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}

// O(N * KlogK)
func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)

	// the O(N) here is N is no of words
	for _, word := range strs {
		// K log K
		sortedWord := SortWord(word)
		mp[sortedWord] = append(mp[sortedWord], word)
	}

	result := make([][]string, 0, len(mp))

	for _, group := range mp {
		result = append(result, group)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
