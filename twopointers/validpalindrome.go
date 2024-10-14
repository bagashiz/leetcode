package twopointers

import "strings"

// https://leetcode.com/problems/valid-palindrome/
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphanum(rune(s[l])) {
			l++
		}

		for r > l && !isAlphanum(rune(s[r])) {
			r--
		}

		if s[l] != s[r] {
			return false
		}

		l, r = l+1, r-1
	}

	return true
}

// isAlphanum checks if a character is alphanumeric
func isAlphanum(r rune) bool {
	return (r >= rune('A') && r <= rune('Z')) ||
		(r >= rune('a') && r <= rune('z')) ||
		(r >= rune('0') && r <= rune('9'))
}
