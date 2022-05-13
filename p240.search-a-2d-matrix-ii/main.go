package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	i := m - 1
	j := 0
	k := m + n - 2

	v := matrix[i][j]
	if v == target {
		return true
	}

	for x := 0; x <= k; x++ {
		v := matrix[i][j]
		if v < target {
			if j == n-1 {
				return false
			}
			j++
		} else if v > target {
			if i == 0 {
				return false
			}
			i--
		} else {
			return true
		}
	}
	return false
}
