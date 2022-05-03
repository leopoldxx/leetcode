package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	sum := make([]int, len(grid[0]))
	sum[0] = grid[0][0]
	for j := 1; j < len(grid[0]); j++ {
		sum[j] = sum[j-1] + grid[0][j]
	}
	for i := 1; i < len(grid); i++ {
		sum[0] += grid[i][0]
		for j := 1; j < len(grid[i]); j++ {
			sum[j] = min(sum[j-1]+grid[i][j], sum[j]+grid[i][j])
		}
	}
	return sum[len(grid[0])-1]
}
