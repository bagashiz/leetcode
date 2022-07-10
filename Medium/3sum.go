package main

import (
	"fmt"
	"sort"
)

// Link: https://leetcode.com/problems/3sum/

func threeSum(nums []int) [][]int {
	// create a 2d array to store the result
	var result [][]int

	// if the length of the array is less than 3, return the result
	if len(nums) < 3 {
		return result
	}

	// sort the array
	sort.Ints(nums)

	// iterate through the array
	for i := 0; i < len(nums)-2; i++ {
		// if the current element is equal to the previous element, continue
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// set the left and right pointers
		l, r := i+1, len(nums)-1

		// iterate through the array
		for l < r {
			// if the sum of the current element and the left and right pointers is equal to 0,
			// add the current element and the left and right pointers to the result
			if nums[i]+nums[l]+nums[r] == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--

				// loop until the left pointer is equal to the right pointer
				for l < r && nums[l] == nums[l-1] {
					l++
				}

				// loop until the right pointer is equal to the left pointer
				for l < r && nums[r] == nums[r+1] {
					r--
				}

			} else if nums[i]+nums[l]+nums[r] < 0 { // if the sum is less than 0, move the left pointer to the right
				l++
			} else {
				r--
			}
		}
	}
	return result
}

func main() {
	// test cases
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1 -1 2], [-1 0 1]]
	fmt.Println(threeSum([]int{0, 1, 1}))             // []
	fmt.Println(threeSum([]int{0, 0, 0}))             // [[0 0 0]]
}
