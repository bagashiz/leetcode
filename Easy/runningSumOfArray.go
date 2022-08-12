package main

import "fmt"

// Link: https://leetcode.com/problems/running-sum-of-1d-array/

func runningSum(nums []int) []int {
	// loop through the array
	for i := 1; i < len(nums); i++ {
		// add the current index value to the previous index value
		nums[i] += nums[i-1]
	}

	return nums
}

func main() {
	// test cases
	fmt.Println(runningSum([]int{1, 2, 3, 4}))     // [1,3,6,10]
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))  // [1,2,3,4,5]
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1})) // [3,4,6,16,17]
}
