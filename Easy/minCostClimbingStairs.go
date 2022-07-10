package main

import (
	"fmt"
	"math"
)

// Link: https://leetcode.com/problems/min-cost-climbing-stairs/

func minCostClimbingStairs(cost []int) int {
	// if the cost length is 0, return 0
	if len(cost) == 0 {
		return 0
	}

	// if the cost length is 1, return the cost[0]
	if len(cost) == 1 {
		return cost[0]
	}

	// if the cost length is 2, compare the cost[0] and cost[1]
	if len(cost) == 2 {
		return int(math.Min(float64(cost[0]), float64(cost[1])))
	}

	// create a copy of the cost array to store the result
	dp := make([]int, len(cost))

	// set the first two elements of the dp array to the cost array
	dp[0] = cost[0]
	dp[1] = cost[1]

	// loop through the cost array from the third element
	for i := 2; i < len(cost); i++ {
		// compare the cost[i] with the sum of the cost[i-1] and cost[i-2]
		dp[i] = int(math.Min(float64(dp[i-1]+cost[i]), float64(dp[i-2]+cost[i])))
	}

	// return the last element of the dp array
	return int(math.Min(float64(dp[len(cost)-1]), float64(dp[len(cost)-2])))
}

func main() {
	// test cases
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))                         // 15
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})) // 6
}
