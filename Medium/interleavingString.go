package main

// Link: https://leetcode.com/problems/interleaving-string/

func isInterleave(s1, s2, s3 string) bool {
	// check if length of s3 is equal to s1+s2
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	// create a 2D array to store the interleaving
	dp := make([][]bool, len(s1)+1)

	// initialize the first row of the 2D array
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	// initialize the first column of the 2D array
	dp[0][0] = true
	for i := 1; i <= len(s1); i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}

	for j := 1; j <= len(s2); j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	// fill the rest of the 2D array
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[len(s1)][len(s2)]
}

func main() {
	// test cases
	println(isInterleave("aabcc", "dbbca", "aadbbcbcac")) // true
	println(isInterleave("aabcc", "dbbca", "aadbbbaccc")) // false
}
