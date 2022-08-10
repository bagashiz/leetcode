package main

import "fmt"

// Link: https://leetcode.com/problems/concatenation-of-array/

func getConcatenation(nums []int) []int {
	// return an array of double size with the original array elements
	return append(nums, nums...)
}

func main() {
	/// test cases
	fmt.Println(getConcatenation([]int{1, 2, 1}))    // [1,2,1,1,2,1]
	fmt.Println(getConcatenation([]int{1, 3, 2, 1})) // [1,3,2,1,1,3,2,1]
}
