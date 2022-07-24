package main

import "fmt"

// Link: https://leetcode.com/problems/move-zeroes/

func moveZeroes(nums []int) {
	// loop through the array
	for i := 0; i < len(nums); i++ {
		// if the current element is 0
		if nums[i] == 0 {
			// loop through the array again
			for j := i + 1; j < len(nums); j++ {
				// if the current element is not 0
				if nums[j] != 0 {
					// swap the current element with the next element
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}

func main() {
	// test cases
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums) // [1, 3, 12, 0, 0]

	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums) // [0]
}
