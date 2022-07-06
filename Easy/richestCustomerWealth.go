package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	// create a slice of ints to hold the sums of each slice
	sum := make([]int, len(accounts))

	// sum the values in each slices of the accounts
	for i := 0; i < len(accounts); i++ {
		for j := 0; j < len(accounts[i]); j++ {
			sum[i] += accounts[i][j]
		}
	}

	// get the max value in the sum slice
	max := sum[0]
	for i := 0; i < len(sum); i++ {
		if sum[i] > max {
			max = sum[i]
		}
	}
	return max
}

func main() {
	// test cases
	accounts := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	accounts2 := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(maximumWealth(accounts))  // 24
	fmt.Println(maximumWealth(accounts2)) // 6
}
