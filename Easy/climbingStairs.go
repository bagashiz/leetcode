package main

// Link: https://leetcode.com/problems/climbing-stairs/

func climbStairs(n int) int {
	// create 2 int variables to store the value of 1
	x, y := 1, 1
	// loop through n-1
	for i := 0; i < n-1; i++ {
		// create temporary variable to store the value of x
		temp := x
		// set x to x + y
		x += y
		// set y to the value of x before x + y
		y = temp
	}
	// return x
	return x
}

func main() {
	// test cases
	println(climbStairs(2)) // 2
	println(climbStairs(3)) // 3
	println(climbStairs(5)) // 8
}
