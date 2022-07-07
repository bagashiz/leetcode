package main

import "fmt"

// Link: https://leetcode.com/problems/ransom-note/

func canConstruct(ransomNote, magazine string) bool {
	// create map consisting of the characters in the magazine
	magazineMap := make(map[rune]int)
	for _, char := range magazine {
		magazineMap[char]++
	}

	// loop through the ransomNote
	for _, char := range ransomNote {
		// if the character is not in the magazineMap, return false
		if magazineMap[char] == 0 {
			return false
		}
		// otherwise, decrement the value of the character in the magazineMap
		magazineMap[char]--
	}

	return true
}

func main() {
	// test cases
	fmt.Println(canConstruct("a", "b"))    // false
	fmt.Println(canConstruct("aa", "ab"))  // false
	fmt.Println(canConstruct("aa", "aab")) // true
}
