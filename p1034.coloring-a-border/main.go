package main

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	return colorBorderWithOldColor(grid, row, col, grid[row][col], color, make(map[[2]int]struct{}))
}

func isBorder(grid [][]int, i int, j int, color int, discovered map[[2]int]struct{}) bool {
	if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[0])-1 {
		return true
	}
	isInArea := func(i, j int) bool {
		_, exists := discovered[[2]int{i, j}]
		return grid[i][j] == color || exists
	}
	if !isInArea(i-1, j) || !isInArea(i+1, j) || !isInArea(i, j-1) || !isInArea(i, j+1) {
		return true
	}
	return false
}

func colorBorderWithOldColor(grid [][]int, row int, col int, oldColor int, newColor int, discovered map[[2]int]struct{}) [][]int {
	_, exists := discovered[[2]int{row, col}]

	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || exists {
		return grid
	}
	if grid[row][col] != oldColor {
		return grid
	}
	discovered[[2]int{row, col}] = struct{}{}
	if isBorder(grid, row, col, oldColor, discovered) {
		grid[row][col] = newColor
	}
	colorBorderWithOldColor(grid, row-1, col, oldColor, newColor, discovered)
	colorBorderWithOldColor(grid, row+1, col, oldColor, newColor, discovered)
	colorBorderWithOldColor(grid, row, col-1, oldColor, newColor, discovered)
	colorBorderWithOldColor(grid, row, col+1, oldColor, newColor, discovered)
	return grid

}
