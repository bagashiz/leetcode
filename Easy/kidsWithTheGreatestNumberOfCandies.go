package main

import "fmt"

// Link: https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	// create a variable to store the max number of candies
	var max int

	// loop through the candies array
	for _, candy := range candies {
		// if the current candy is greater than the max
		if candy > max {
			// set the max to the current candy
			max = candy
		}
	}

	// create a slice of booleans with the same length as the candies array to store the results
	res := make([]bool, len(candies))

	// loop through the candies array
	for i, candy := range candies {
		// set the current index of the results array to true if the current candy plus the extra candies is greater than or equal to the max
		res[i] = candy+extraCandies >= max
	}

	return res
}

func main() {
	// test cases
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)) // [true true true false true]
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1)) // [true false false false false]
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))    // [true false true]
}
