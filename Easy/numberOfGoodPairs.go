package main

// Link: https://leetcode.com/problems/number-of-good-pairs/

func numIdenticalPairs(nums []int) int {
	// create a variable to store the number of good pairs
	var res int

	// loop through the array
	for i := 0; i < len(nums); i++ {
		// loop through the array again to compare the current value with the rest of the array
		for j := i + 1; j < len(nums); j++ {
			// if the current value is equal to the next value
			if nums[i] == nums[j] {
				// increment the number of good pairs
				res++
			}
		}
	}
	return res
}

func main() {
	// test cases
	println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})) // 4
	println(numIdenticalPairs([]int{1, 1, 1, 1}))       // 6
	println(numIdenticalPairs([]int{1, 2, 3}))          // 0
}
