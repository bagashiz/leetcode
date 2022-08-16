package main

// Link: https://leetcode.com/problems/first-unique-character-in-a-string/

func firstUniqChar(s string) int {
	// create a map of integer to store the count of each alphabet character in the string
	var charMap [26]int

	// loop through the string
	for _, c := range s {
		// increment the count of each character in the map
		charMap[c-'a']++
	}

	// loop through the string again
	for i, c := range s {
		// if the count of the character is 1, return the index
		if charMap[c-'a'] == 1 {
			return i
		}
	}

	// if no unique character found, return -1
	return -1
}

func main() {
	// test cases
	println(firstUniqChar("leetcode"))     // 0
	println(firstUniqChar("loveleetcode")) // 2
	println(firstUniqChar("aabb"))         // -1
}
