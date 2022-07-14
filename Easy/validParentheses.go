package main

// Link: https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	// create a slice to store the parentheses
	var parentheses []rune

	// create a map to store all known parentheses
	var parenthesesMap = map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// iterate through the string
	for _, r := range s {
		// switch on the rune
		switch r {
		// if the rune is a opening parenthesis
		case '(', '{', '[':
			// append the rune to the slice
			parentheses = append(parentheses, r)
		// if the rune is a closing parenthesis
		case ')', '}', ']':
			// if the slice is empty, or the last element is not a matching parenthesis
			if len(parentheses) == 0 || parentheses[len(parentheses)-1] != parenthesesMap[r] {
				// return false
				return false
			}

			// remove the last element from the slice
			parentheses = parentheses[:len(parentheses)-1]
		}
	}

	// if the slice is empty, return true
	return len(parentheses) == 0
}

func main() {
	// test cases
	println(isValid("()"))     // true
	println(isValid("()[]{}")) // true
	println(isValid("(]"))     // false
	println(isValid("([)]"))   // false
	println(isValid("([])"))   // true
}
