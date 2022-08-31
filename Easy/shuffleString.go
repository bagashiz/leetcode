package main

// Link: https://leetcode.com/problems/shuffle-string/

func restoreString(s string, indices []int) string {
	// create an array to store the answer
	res := make([]byte, len(s))

	// loop through indices
	for index, value := range indices {
		// store the character at the index in the answer array
		res[value] = s[index]
	}

	// return the converted res array to string
	return string(res)
}

func main() {
	// test cases
	println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 1, 2, 3})) // "leetcode"
	println(restoreString("abc", []int{0, 1, 2}))                     // "abc"
}
