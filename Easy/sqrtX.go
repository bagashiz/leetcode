package main

// Link: https://leetcode.com/problems/sqrtx/

func mySqrt(x int) int {
	// loop until x
	for i := 0; i <= x; i++ {
		// if i^2 is equal to x
		if i*i == x {
			return i
		}

		// if i^2 is greater than x
		if i*i > x {
			return i - 1
		}
	}

	return 0
}

func main() {
	// test cases
	println(mySqrt(4))  // 2
	println(mySqrt(8))  // 2
	println(mySqrt(10)) // 3
}
