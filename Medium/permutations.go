package main

import "fmt"

// Link: https://leetcode.com/problems/permutations/

func permute(nums []int) [][]int {
	// if length of nums is 1, return nums
	if len(nums) == 1 {
		return [][]int{nums}
	}

	// create variable to store permutations
	result := [][]int{}

	// loop number in nums
	for i := 0; i < len(nums); i++ {
		// create variable to store copyNums of nums
		copyNums := make([]int, len(nums))
		copy(copyNums, nums)

		// swap nums[i] with nums[0]
		copyNums[i], copyNums[0] = copyNums[0], copyNums[i]

		// remove nums[i] from copyNums
		copyNums = copyNums[1:]

		// create variable to store permutations of copyNums
		permutations := permute(copyNums)

		// loop permutations of copyNums
		for _, permutation := range permutations {
			// append nums[i] to each permutation
			permutation = append([]int{nums[i]}, permutation...)
			// append permutation to result
			result = append(result, permutation)
		}
	}

	return result
}

func main() {
	// test cases
	fmt.Println(permute([]int{1, 2, 3})) // [ [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1] ]
	fmt.Println(permute([]int{0, 1}))    // [ [0,1], [1,0] ]
	fmt.Println(permute([]int{1}))       // [ [1] ]
}
