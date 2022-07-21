package main

import "fmt"

// Link: https://leetcode.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
	// if digits is empty, return empty array of strings
	if len(digits) == 0 {
		return []string{}
	}

	// create map of digits to letters
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	// create an array of strings to store results
	var res []string

	// create a function to recursively generate combinations
	var generate func(string, int)
	generate = func(current string, index int) {
		// if index is out of bounds, append current to res and return
		if index >= len(digits) {
			res = append(res, current)
			return
		}

		// get the letters for the current digit
		letters := m[digits[index]]

		// for each letter, append it to current and call generate again
		for _, letter := range letters {
			generate(current+letter, index+1)
		}
	}

	// call generate with empty string and 0 to start recursion and append results to res
	generate("", 0)

	// return res
	return res
}

func main() {
	// test cases
	fmt.Println(letterCombinations("23")) // ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
	fmt.Println(letterCombinations(""))   // []
	fmt.Println(letterCombinations("2"))  // ["a", "b", "c"]
}
