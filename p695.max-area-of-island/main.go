package main

func count(grid [][]int, x, y int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
		return 0
	}
	grid[x][y] = 0
	return 1 + count(grid, x-1, y) + count(grid, x+1, y) + count(grid, x, y-1) + count(grid, x, y+1)
}

func maxAreaOfIsland(grid [][]int) int {
	maxCount := 0

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 1 {
				currentCount := count(grid, x, y)
				if currentCount > maxCount {
					maxCount = currentCount
				}
			}
		}
	}
	return maxCount
}
