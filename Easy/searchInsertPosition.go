package main

// Link: https://leetcode.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	// loop through the array
	for i, v := range nums {
		// if the target is less than the current value, return the index
		if v >= target {
			return i
		}
	}

	// if the target is greater than all the values, return the length of the array
	// because the length of the array is the index of the last value + 1
	return len(nums)
}

func main() {
	// test cases
	println(searchInsert([]int{1, 3, 5, 6}, 5)) // 2
	println(searchInsert([]int{1, 3, 5, 6}, 2)) // 1
	println(searchInsert([]int{1, 3, 5, 6}, 7)) // 4
}
