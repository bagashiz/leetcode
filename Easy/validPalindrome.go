package main

import "strings"

// Link: https://leetcode.com/problems/valid-palindrome/

func isPalindrome(s string) bool {
	// create two variables for pointers
	i, j := 0, len(s)-1

	// loop until i = j
	for i < j {
		// if s[i] is not an alphanumeric character, increment i and continue
		if !isValid(s[i]) {
			i++
			continue
		}

		// if s[j] is not an alphanumeric character, decrement j and continue
		if !isValid(s[j]) {
			j--
			continue
		}

		// if s[i] is not equal to s[j], return false
		if !strings.EqualFold(string(s[i]), string(s[j])) {
			return false
		}
		// increment i and decrement j
		i++
		j--
	}
	return true
}

// isValid checks if a character is an alphanumeric character
func isValid(a byte) bool {
	// return true if a contains a-z, A-Z, or 0-9
	if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

func main() {
	// test cases
	println(isPalindrome("A man, a plan, a canal: Panama")) // true
	println(isPalindrome("race a car"))                     // false
	println(isPalindrome("race car"))                       // true
	println(isPalindrome(" "))                              // true
}
