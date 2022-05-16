package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := [101]int{}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	num := 1
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			num = 0
		}
		dp[j] = num
	}

	for i := 1; i < m; i++ {

		next := 0
		for k := 0; k < n; k++ {
			if dp[k] == 0 {
				next++
			} else {
				break
			}
		}
		if next == n {
			return 0
		}
		if obstacleGrid[i][next] == 1 {
			dp[next] = 0
		}
		for j := next + 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			dp[j] = dp[j-1] + dp[j]
		}
	}
	return dp[n-1]

}
