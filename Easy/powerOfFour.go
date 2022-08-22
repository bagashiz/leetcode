package main

// Link: https://leetcode.com/problems/power-of-four/

func isPowerOfFour(n int) bool {
	// if n is 0 or negative, return false
	if n <= 0 {
		return false
	}

	// while n is divisible by 4
	for n%4 == 0 {
		// divide it by 4
		n /= 4
	}

	// return true if n is 1, false otherwise
	return n == 1
}

func main() {
	// test cases
	println(isPowerOfFour(16)) // true
	println(isPowerOfFour(5))  // false
	println(isPowerOfFour(1))  // true
}
