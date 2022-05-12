package main

func min(a, b, c int16) int16 {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}
func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	max := 0

	area := make([][]int16, 0, m)
	area = append(area, make([]int16, n))
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			area[0][j] = 1
			max = 1
		}
	}

	for i := 1; i < m; i++ {
		area = append(area, make([]int16, n))
		if matrix[i][0] == '1' {
			area[i][0] = 1
			if max == 0 {
				max = 1
			}
		}

		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				area[i][j] = min(area[i-1][j], area[i][j-1], area[i-1][j-1]) + 1
				if max < int(area[i][j]) {
					max = int(area[i][j])
				}
			}
		}
	}
	return max * max
}
