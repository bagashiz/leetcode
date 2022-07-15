package main

import "unicode"

// Link: https://leetcode.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	// if s is empty, return 0
	if len(s) == 0 {
		return 0
	}

	// loop through s to remove all spaces at the end of the string
	l := len(s)
	for l > 0 && unicode.IsSpace(rune(s[l-1])) {
		// decrement l
		l--
		// trim the string to the last non-space character
		s = s[:l]
	}

	// count the length of the last word
	length := 0
	for i := l - 1; i >= 0; i-- {
		// if the rune is not a space, increment length
		if !unicode.IsSpace(rune(s[i])) {
			length++
		} else {
			// if the rune is a space, break
			break
		}
	}

	return length
}

func main() {
	// test cases
	println(lengthOfLastWord("Hello World"))                 // 5
	println(lengthOfLastWord("luffy is still joyboy"))       // 6
	println(lengthOfLastWord("   fly me   to   the moon  ")) // 4
}
