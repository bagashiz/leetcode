package arrayshashing

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/valid-anagram/
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s = sortString(s)
	t = sortString(t)
	return s == t
}

// sortString rearrange characters in a string alphabetically
func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
