package main

import "math"

// Link: https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	// if string is empty, return 0
	if len(s) == 0 {
		return 0
	}

	// create a variable to store the longest length
	longest := 0

	// create a variable to store the current length
	current := 0

	// create a map to store the char and its index
	m := make(map[rune]int)

	// loop
	for i := 0; i < len(s); i++ {
		// if the char is already in the map, update the current
		if _, ok := m[rune(s[i])]; ok {
			// update the current length to the index of the char + 1
			current = int(math.Max(float64(current), float64(m[rune(s[i])]+1)))
		}

		// update the map with the char and its index
		m[rune(s[i])] = i

		// update the longest length to the current length
		longest = int(math.Max(float64(longest), float64(i-current+1)))
	}

	return longest
}

func main() {
	//test cases
	println(lengthOfLongestSubstring("abcabcbb")) // 3
	println(lengthOfLongestSubstring("bbbbb"))    // 1
	println(lengthOfLongestSubstring("pwwkew"))   // 3
}
