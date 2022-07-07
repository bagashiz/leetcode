package main

// Link: https://leetcode.com/problems/two-sum/

import "fmt"

func twoSum(nums []int, target int) []int {
	// create a map to store the values
	m := make(map[int]int)

	// iterate through the array
	for i, v := range nums {
		// if the target - v is in the map, return the indices
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		// otherwise, add the value to the map
		m[v] = i
	}
	// if the target is not in the map, return nil
	return nil
}

func main() {
	// test cases
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target)) // [0, 1]

	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(nums, target)) // [1, 2]

	nums = []int{3, 3}
	target = 6
	fmt.Println(twoSum(nums, target)) // [0, 1]
}
