package main

// Link: https://leetcode.com/problems/single-number/

func singleNumber(nums []int) int {
	// create a map to store the number of each number
	numMap := make(map[int]int)

	// loop through the array
	for _, num := range nums {
		// increment the number of the current number
		numMap[num]++
	}

	// loop through the map
	for num, count := range numMap {
		// if the number of the current number is 1, return it
		if count == 1 {
			return num
		}
	}

	// return 0 if no number is found
	return 0
}

func main() {
	// test cases
	println(singleNumber([]int{2, 2, 1}))       // 1
	println(singleNumber([]int{4, 1, 2, 1, 2})) // 4
	println(singleNumber([]int{1}))             // 1
}
