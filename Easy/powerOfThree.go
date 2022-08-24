package main

// Link: https://leetcode.com/problems/power-of-three/

func isPowerOfThree(n int) bool {
	// if n is 0, return false
	if n == 0 {
		return false
	}

	// while the remainder of n / 3 is 0
	for n%3 == 0 {
		// divide n by 3
		n /= 3
	}

	// if n is 1, return true
	return n == 1
}

func main() {
	// test cases
	println(isPowerOfThree(27)) // true
	println(isPowerOfThree(0))  // false
	println(isPowerOfThree(9))  // true
}
