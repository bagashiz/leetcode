package main

import (
	"math"
	"strings"
)

// Link: https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	// if s is empty or only contains spaces, return 0
	if len(s) == 0 || strings.TrimSpace(s) == "" {
		return 0
	}

	// remove leading spaces
	s = strings.TrimSpace(s)

	// create a variable to store index of the string
	i := 0
	// create sign variable to store sign of number
	var sign int
	// set sign to 1 if s[i] is '+' or -1 if s[i] is '-'
	switch s[i] {
	case '+':
		// set sign to 1
		sign = 1
		// skip to the next character
		i++
	case '-':
		// set sign to -1
		sign = -1
		// skip to the next character
		i++
	default:
		sign = 1
	}

	// create a variable to store result
	var res int

	// loop through s and check if s[i] is a digit
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		// convert s[i] to int and add it to res
		res = res*10 + int(s[i]-'0')

		if res*sign > math.MaxInt32 {
			// if res is greater than 2147483647, return 2147483647
			return math.MaxInt32

		} else if res*sign < math.MinInt32 {
			// if res is less than -2147483648, return -2147483648
			return math.MinInt32
		}
		// skip to next character
		i++
	}

	// return res multiplied by sign (positive or negative)
	return res * sign
}

func main() {
	// test cases
	println(myAtoi(" "))               // 0
	println(myAtoi("42"))              // 42
	println(myAtoi("   -42"))          // -42
	println(myAtoi("4193 with words")) // 4193
	println(myAtoi("words and 987"))   // 0
}
