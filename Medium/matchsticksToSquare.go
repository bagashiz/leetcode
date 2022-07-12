package main

import "fmt"

// Link: https://leetcode.com/problems/matchsticks-to-square/

func makeSquare(matchsticks []int) bool {
	// if the length of the matchsticks is less than 4, return false
	if len(matchsticks) < 4 {
		return false
	}

	// check if the sum of the matchsticks is a square
	sum := 0
	for _, v := range matchsticks {
		sum += v
	}
	// if the sum is not divisible by 4, then it is not a square
	if sum%4 != 0 {
		return false
	}

	return findCombinations(matchsticks, sum, 0, 0, 0, 0, 0)
}

// function to find combinations of matchsticks
func findCombinations(nums []int, sum, idx, s1, s2, s3, s4 int) bool {
	// check if the index is equal to the length of the matchsticks
	if idx == len(nums) {
		// if s1 is equal to s2, s2 is equal to s3, s3 is equal to s4, and s4 is equal to s1, then return true
		if s1 == s2 && s2 == s3 && s3 == s4 && s4 == s1 {
			return true
		}
		// else return false
		return false
	}

	// check if nums[idx] + s1 is less than or equal to the sum divided by 4
	if nums[idx]+s1 <= sum/4 {
		// if the function is called again with the index incremented, s1 is incremented by nums[idx] and the function is called again
		if findCombinations(nums, sum, idx+1, s1+nums[idx], s2, s3, s4) {
			return true
		}
	}

	// check if nums[idx] + s2 is less than or equal to the sum divided by 4
	if nums[idx]+s2 <= sum/4 {
		// if the function is called again with the index incremented, s2 is incremented by nums[idx] and the function is called again
		if findCombinations(nums, sum, idx+1, s1, s2+nums[idx], s3, s4) {
			return true
		}
	}

	// check if nums[idx] + s3 is less than or equal to the sum divided by 4
	if nums[idx]+s3 <= sum/4 {
		// if the function is called again with the index incremented, s3 is incremented by nums[idx] and the function is called again
		if findCombinations(nums, sum, idx+1, s1, s2, s3+nums[idx], s4) {
			return true
		}
	}

	// check if nums[idx] + s4 is less than or equal to the sum divided by 4
	if nums[idx]+s4 <= sum/4 {
		// if the function is called again with the index incremented, s4 is incremented by nums[idx] and the function is called again
		if findCombinations(nums, sum, idx+1, s1, s2, s3, s4+nums[idx]) {
			return true
		}
	}

	// if none of the above conditions are true, return false
	return false
}

func main() {
	// test cases
	fmt.Println(makeSquare([]int{}))                                                       // false
	fmt.Println(makeSquare([]int{1, 1, 2, 2, 2}))                                          // true
	fmt.Println(makeSquare([]int{3, 3, 3, 3, 4}))                                          // false
	fmt.Println(makeSquare([]int{9, 38, 1, 76, 41, 19, 3, 97, 48, 5, 83, 15, 51, 99, 95})) // false
}
