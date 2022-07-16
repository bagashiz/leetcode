package main

import "fmt"

// Link: https://leetcode.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	// loop through the array
	for i := 0; i < len(nums); i++ {
		// if the current element is equal to the value
		if nums[i] == val {
			// remove the element by shifting the array to the left
			nums = append(nums[:i], nums[i+1:]...)
			// decrement the index by 1 to account for the removed element
			// and move back to the current index to check again if it is equal to the value
			i--
		}
	}
	// return the length of the array
	return len(nums)
}

func main() {
	// test cases
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))             // 2
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)) // 5
}
