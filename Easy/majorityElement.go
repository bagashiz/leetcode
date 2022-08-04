package main

// Link: https://leetcode.com/problems/majority-element/

func majorityElement(nums []int) int {
	// create a map to store the count of each element
	m := make(map[int]int)

	// loop through the array
	for _, c := range nums {
		// if the element is already in the map, increment the count
		m[c]++
	}

	// loop through the map
	for e, c := range m {
		// if the count is greater than half the length of the array, return the element
		if c > len(nums)/2 {
			return e
		}
	}

	// if the element is not in the map, return 0
	return 0
}

func main() {
	// test cases
	println(majorityElement([]int{3, 2, 3}))             // 3
	println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2})) // 2
}
