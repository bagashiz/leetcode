package main

import (
	"fmt"
	"sort"
)

// Link: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

// function to find the index of an element in a slice
func indexOf(element int, data []int) int {
	for i, e := range data {
		if element == e {
			return i
		}
	}
	return -1 //not found.
}

func kWeakestRows(mat [][]int, k int) []int {
	// variable to store the weakest rows
	var weakestRows []int

	// loop through the matrix
	for i := 0; i < len(mat); i++ {
		// variable to store the sum of the row
		var sum int

		// loop through the row
		for j := 0; j < len(mat[i]); j++ {
			// add the value to the sum
			sum += mat[i][j]
		}

		// add the row to the weakest rows
		weakestRows = append(weakestRows, sum)
	}

	// make a copy of the weakest rows then sort it
	weakestRowsSorted := append([]int{}, weakestRows...)
	sort.Ints(weakestRowsSorted)

	// find the index of the weakest rows
	var weakestRowsIndex []int
	for i := 0; i < k; i++ {
		index := indexOf(weakestRowsSorted[i], weakestRows) // find the index of the weakest row
		weakestRows[index] = -1                             // set the weakest row to -1 so it won't be added to the result again
		weakestRowsIndex = append(weakestRowsIndex, index)  // add the index to the weakestRowsIndex
	}
	// return the weakest rows index
	return weakestRowsIndex
}

func main() {
	// test cases
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}, 3)) // [2 0 3]

	fmt.Println(kWeakestRows([][]int{
		{0, 1, 1},
		{1, 1, 1},
		{0, 0, 0},
	}, 2)) // [2 0]

	fmt.Println(kWeakestRows([][]int{
		{0, 0, 1, 1, 0},
		{0, 0, 1, 1, 0},
		{0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
	}, 1)) // [3]
}
