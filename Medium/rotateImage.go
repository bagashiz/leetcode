package main

import "fmt"

// Link: https://leetcode.com/problems/rotate-image/

func rotate(matrix [][]int) {
	// create a variable to store the length of the matrix
	n := len(matrix)

	// loop through half of the matrix length
	for i := 0; i < n/2; i++ {
		// loop again through current matrix element until n-i-1
		for j := i; j < n-i-1; j++ {
			// rotate the current matrix row
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] = matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
		}
	}
}

func main() {
	// test cases
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrix2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	rotate(matrix1)
	rotate(matrix2)

	fmt.Println(matrix1) // [[7 4 1] [8 5 2] [9 6 3]]
	fmt.Println(matrix2) // [[15 13 2 4] [14 3 5 1] [12 6 8 9] [16 7 10 11]]
}
