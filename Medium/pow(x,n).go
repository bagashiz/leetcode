package main

import (
	"fmt"
	"math"
)

// Link: https://leetcode.com/problems/powx-n/

func myPow(x float64, n int) float64 {
	// call math.Pow() function to calculate the power of x to n and return the result
	return math.Pow(x, float64(n))
}

func main() {
	// test cases
	fmt.Println(myPow(2, 10))  // 1024.00000
	fmt.Println(myPow(2.1, 3)) // 9.26100
	fmt.Println(myPow(2, -2))  // 0.25000
}
