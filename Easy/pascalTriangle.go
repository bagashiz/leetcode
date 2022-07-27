package main

import "fmt"

// Link: https://leetcode.com/problems/pascals-triangle/

func generate(numRows int) [][]int {
	// if numRows equals to 0, return empty array
	if numRows == 0 {
		return nil
	}

	// if numRows equals to 1, return array with value [1]
	if numRows == 1 {
		return [][]int{{1}}
	}

	// if numRows equals to 2, return array with value [1] and [1, 1]
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	// create an array to store result that already contains [1] and [1, 1]
	result := [][]int{{1}, {1, 1}}

	// loop from rows 3 to numRows
	for i := 2; i < numRows; i++ {
		// create a temporary array to store first column which is always 1 for the first column
		temp := []int{1}

		// loop from index 0 to i - 1
		for j := 0; j < i-1; j++ {
			// add current row's value to temp array using pascal's formula
			temp = append(temp, result[i-1][j]+result[i-1][j+1])
		}

		// add 1 to the last element of temp array
		temp = append(temp, 1)

		// add temp array to result array
		result = append(result, temp)
	}
	return result
}

func main() {
	// test cases
	fmt.Println(generate(1))  // [[1]]
	fmt.Println(generate(5))  // [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1]]
	fmt.Println(generate(10)) // [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1], [1, 5, 10, 10, 5, 1], [1, 6, 15, 20, 15, 6, 1], [1, 7, 21, 35, 35, 21, 7, 1], [1, 8, 28, 56, 70, 56, 28, 8, 1], [1, 9, 36, 84, 126, 126, 84, 36, 9, 1]]
}
