package main

func uniquePaths(m int, n int) int {
	dp := [101]int{}
	if m == 1 || n == 1 {
		return 1
	}
	for j := 0; j < n; j++ {
		dp[j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}
	return dp[n-1]
}
