package main

// Link: https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/

func numberOfSteps(num int) int {
	// variable to store the iterations
	var steps int

	// loop until the number is zero
	for num != 0 {
		// if the number is even, divide by 2
		if num%2 == 0 {
			num /= 2 // divide by 2
		} else {
			num-- // subtract 1
		}
		steps++ // increment the steps
	}
	return steps
}

func main() {
	// test cases
	println(numberOfSteps(14))  // 6
	println(numberOfSteps(99))  // 10
	println(numberOfSteps(69))  // 9
	println(numberOfSteps(420)) // 12
}
