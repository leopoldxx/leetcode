package main

func sink(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'

	sink(grid, i-1, j)
	sink(grid, i+1, j)
	sink(grid, i, j-1)
	sink(grid, i, j+1)
}

func numIslands(grid [][]byte) int {
	total := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				sink(grid, i, j)
				total++
			}
		}
	}
	return total
}
