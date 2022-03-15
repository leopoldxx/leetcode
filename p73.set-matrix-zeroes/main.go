package main

func findFirstZero(matrix [][]int) (bool, int, int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				return true, i, j
			}
		}
	}
	return false, -1, -1
}

func setZeroes(matrix [][]int) {
	found, x, y := findFirstZero(matrix)
	if !found {
		return
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][y] = 0
				matrix[x][j] = 0
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == x || j == y {
				continue
			}
			if matrix[x][j] == 0 || matrix[i][y] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		matrix[i][y] = 0
	}
	for j := 0; j < len(matrix[0]); j++ {
		matrix[x][j] = 0
	}

}
