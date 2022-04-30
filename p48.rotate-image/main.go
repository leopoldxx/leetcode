package main

func rotate(matrix [][]int) {
	i := 0
	j := len(matrix) - 1
	for i < j {
		matrix[i], matrix[j] = matrix[j], matrix[i]
		i++
		j--
	}
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
