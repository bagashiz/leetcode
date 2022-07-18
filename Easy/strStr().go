package main

// Link: https://leetcode.com/problems/implement-strstr/

func strstr(haystack string, needle string) int {
	// check if needle is empty
	if needle == "" {
		return 0
	}

	// loop through the length of haystack minus the length of needle + 1
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		// check if the characters at the current index of haystack match the first character of needle
		if haystack[i:i+len(needle)] == needle {
			// return the index of character that matched the character at the current index of haystack
			return i
		}
	}

	// return -1 if no match was found
	return -1
}

func main() {
	// test cases
	println(strstr("hello", "ll"))  // 2
	println(strstr("aaaaa", "bba")) // -1
}
