package main

// Link: https://leetcode.com/problems/reverse-integer/

func reverse(x int) int {
	// create a new int to store the reversed number
	var result int

	// loop through the number and add each digit to the result
	for x != 0 {
		// get the last digit, and remove it from the number
		result = result*10 + x%10
		x = x / 10
	}

	// if the result is greater than 2^31 - 1 or less than -2^31, return 0
	if result > 2147483647 || result < -2147483648 {
		return 0
	}

	// return the reversed number
	return result
}

func main() {
	// test cases
	println(reverse(123))        // 321
	println(reverse(-123))       // -321
	println(reverse(120))        // 21
	println(reverse(1534236469)) // 0
}
