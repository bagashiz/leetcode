package main

import (
	"fmt"
	"sort"
)

// Link: https://leetcode.com/problems/4sum/

func fourSum(nums []int, target int) [][]int {
	// create a 2d array to store the result
	var result [][]int

	// if the length of the array is less than 4, return the result
	if len(nums) < 4 {
		return result
	}

	// sort the array
	sort.Ints(nums)

	// iterate through the array
	for i := 0; i < len(nums)-3; i++ {
		// if the current element is equal to the previous element, continue
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// iterate through the array from i+1 to len(nums)-2
		for j := i + 1; j < len(nums)-2; j++ {
			// if the current element is equal to the previous element, continue
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// set the left and right pointers
			l, r := j+1, len(nums)-1

			// iterate through the array if l < r
			for l < r {
				// if the sum of the current element and the left and right pointers is equal to target,
				// add the current element and the left and right pointers to the result
				if nums[i]+nums[j]+nums[l]+nums[r] == target {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
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

				} else if nums[i]+nums[j]+nums[l]+nums[r] < target { // if the sum is less than target, move the left pointer to the right
					l++
				} else {
					r--
				}
			}
		}
	}

	return result
}

func main() {
	// test cases
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0)) // [[-2 -1 1 2], [-2 0 0 2], [-1 0 0 1]]
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))      // [[2 2 2 2 2]]
}
