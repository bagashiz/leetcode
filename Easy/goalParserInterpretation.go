package main

// Link: https://leetcode.com/problems/goal-parser-interpretation/

func interpret(command string) string {
	// create a variable to store the result
	var result string

	// loop through the command string
	for i := 0; i < len(command); {
		// if the character(s) at the current index to index+1 is "G"
		if command[i:i+1] == "G" {
			// append "G" to the result
			result += "G"
			// increment the index by 1
			i++

			// if the character(s) at the current index to index+2 is "()"
		} else if command[i:i+2] == "()" {
			// append "o" to the result
			result += "o"
			// increment the index by 2
			i += 2

			// else the character(s) must be "(al)"
		} else {
			// append "al" to the result
			result += "al"
			// increment the index by 4
			i += 4
		}
	}

	return result
}

func main() {
	// test cases
	println(interpret("G()(al)"))        // "Goal"
	println(interpret("G()()()()(al)"))  // "Gooooal"
	println(interpret("(al)G(al)()()G")) // "alGalooG"
}
