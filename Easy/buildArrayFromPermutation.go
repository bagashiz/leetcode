package main

import "fmt"

// Link: https://leetcode.com/problems/build-array-from-permutation/

func buildArray(nums []int) []int {
	// create a new array to store the result
	var result []int

	// loop through the array nums
	for _, num := range nums {
		// append the value of the index of the current num value to the result array
		result = append(result, nums[num])
	}

	return result
}

func main() {
	// test cases
	fmt.Println(buildArray([]int{0, 2, 1, 5, 3, 4})) // [0,1,2,4,5,3]
	fmt.Println(buildArray([]int{5, 0, 1, 2, 3, 4})) // [4,5,0,1,2,3]
}
