package main

// Link: https://leetcode.com/problems/final-value-of-variable-after-performing-operations/

func finalValueAfterOperations(operations []string) int {
	// create an integer variable to store the final value with initial value 0
	var x int

	// loop through the operations
	for _, op := range operations {
		// if the operation is ++X or X++ then increment the value of x by 1
		if op[1] == '+' {
			x++
		} else { // else decrement the value of x by 1
			x--
		}
	}

	return x
}

func main() {
	// test cases
	println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))        // 1
	println(finalValueAfterOperations([]string{"++X", "++X", "X++"}))        // 3
	println(finalValueAfterOperations([]string{"X++", "++X", "--X", "X--"})) // 0
}
