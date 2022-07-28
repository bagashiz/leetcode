package main

// Link: https://leetcode.com/problems/valid-anagram/

func isAnagram(s, t string) bool {
	// if the length of s and t are not equal, return false
	if len(s) != len(t) {
		return false
	}

	// create a map to store the characters of s
	m := make(map[rune]int)

	// iterate through the characters of s
	for _, v := range s {
		// increment the value of the key in the map
		m[v]++
	}

	// iterate through the characters of t
	for _, v := range t {
		// decrement the value of the key in the map
		m[v]--
	}

	// iterate through the map
	for _, v := range m {
		// if the value of the key is not 0, return false
		// because the characters of s and t are not anagrams
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	// test cases
	println(isAnagram("anagram", "nagaram")) // true
	println(isAnagram("rat", "car"))         // false
	println(isAnagram("", ""))               // true
	println(isAnagram(" ", ""))              // false
}
