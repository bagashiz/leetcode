package main

import "fmt"

// Link: https://leetcode.com/problems/numbers-smaller-than-current/

func smallerNumbersThanCurrent(nums []int) []int {
	// create an array to store the result
	res := make([]int, len(nums))

	// loop through the array
	for i := 0; i < len(nums); i++ {
		// create a counter to count the number of smaller numbers
		count := 0

		// loop through the array again
		for j := 0; j < len(nums); j++ {
			// if the current number is smaller than the first loop number
			if nums[j] < nums[i] {
				// increment the counter
				count++
			}
		}

		// store the counter in the result array
		res[i] = count
	}

	return res
}

func main() {
	// test cases
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3})) // [4,0,1,1,3]
	fmt.Println(smallerNumbersThanCurrent([]int{6, 5, 4, 8}))    // [2,1,0,3]
	fmt.Println(smallerNumbersThanCurrent([]int{7, 7, 7, 7}))    // [0,0,0,0]
}
