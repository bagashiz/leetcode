package main

import "strings"

// Link: https://leetcode.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	// if there is no string, return empty string
	if len(strs) == 0 {
		return ""
	}

	// create a variable to store the longest common prefix, set it to the first string
	prefix := strs[0]

	// loop through the rest of the strings
	for i := 1; i < len(strs); i++ {
		// if the prefix is not a prefix of the current string
		for strings.Index(strs[i], prefix) != 0 {
			// remove the last character from the prefix
			prefix = prefix[:len(prefix)-1]
		}
	}

	// return the longest common prefix
	return prefix
}

func main() {
	// test  cases
	println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
	println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
}
