package main

// Link: https://leetcode.com/problems/number-of-matching-subsequences/

func numMatchingSubseq(s string, words []string) int {
	// create a variable to store the times of matching subsequences
	var count int

	// loop through the words
	for _, word := range words {
		// create variables to store the index of the word
		var i, j int

		// loop until i is equal to the length of s and j is equal to the length of word
		for i < len(s) && j < len(word) {
			// if the character at index i of s is equal to the character at index j of word
			if s[i] == word[j] {
				// increment i and j
				i++
				j++
			} else {
				// increment i
				i++
			}
		}
		// if the word is a subsequence of s
		if j == len(word) {
			// increment the count
			count++
		}
	}
	return count
}

func main() {
	// test cases
	println(numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))                          // 3
	println(numMatchingSubseq("dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"})) // 2
}
