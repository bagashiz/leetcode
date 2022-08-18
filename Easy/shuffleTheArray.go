package main

import "fmt"

// Link: https://leetcode.com/problems/shuffle-the-array/

func shuffle(nums []int, n int) []int {
	// create a variable to store the result
	var res []int

	// loop through the array until the index is less than n
	for i := 0; i < n; i++ {
		// append the value of the first half of the array to the result
		res = append(res, nums[i])

		// append the value of the second half of the array to the result
		res = append(res, nums[i+n])
	}

	return res
}

func main() {
	// test cases
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))       // [2,3,5,4,1,7]
	fmt.Println(shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4)) // [1,4,2,3,3,2,4,1]
	fmt.Println(shuffle([]int{1, 1, 2, 2}, 2))             // [1,2,1,2]
}
