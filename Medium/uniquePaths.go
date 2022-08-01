package main

// Link: https://leetcode.com/problems/unique-paths/

func uniquePaths(m int, n int) int {
	// create a slice with the capacity of n, and fill the first index with 1
	dp := make([]int, n)
	dp[0] = 1

	// loop through the first row
	for i := 0; i < m; i++ {
		// loop through the second row
		for j := 1; j < n; j++ {
			// add the value of the previous row + the value of the previous column
			dp[j] += dp[j-1]
		}
	}

	// return the last value in the slice
	return dp[n-1]
}

func main() {
	println(uniquePaths(3, 7))  // 28
	println(uniquePaths(3, 2))  // 3
	println(uniquePaths(51, 9)) // 3
}
