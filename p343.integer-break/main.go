package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = i - 1
		for j := i - 2; j > 1; j-- {
			if i-j > dp[i-j] {
				dp[i-j] = i - j
			}
			dp[i] = max(dp[i], dp[i-j]*j)
		}
	}
	return dp[n]
}
