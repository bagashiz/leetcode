package main

import "strings"

// Link: https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/

func mostWordsFound(sentences []string) int {
	// create a variable to store the max number of words
	max := 0

	// loop through the sentences
	for _, sentence := range sentences {
		// split the sentence into words using strings.Fields function
		words := strings.Fields(sentence)

		// if the number of words is greater than the max, update the max
		if len(words) > max {
			max = len(words)
		}
	}

	return max
}

func main() {
	// test cases
	println(mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"})) // 6
	println(mostWordsFound([]string{"please wait", "continue to fight", "continue to win"}))                             // 3
}
