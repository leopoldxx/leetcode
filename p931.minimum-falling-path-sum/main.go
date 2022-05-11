package main

func min3(a, b, c int) int {
	return min2(a, min2(b, c))
}
func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min(line []int, sum []int) []int {
	n := len(line)
	for i := 0; i < n; i++ {
		if i == 0 && i+1 < n {
			line[i] += min2(sum[i], sum[i+1])
		} else if i == n-1 && i-1 >= 0 {
			line[i] += min2(sum[i], sum[i-1])
		} else if n >= 3 {
			line[i] += min3(sum[i-1], sum[i], sum[i+1])
		} else {
			line[i] += sum[i]
		}
	}
	return line
}

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix[0])
	minSum := make([]int, n)
	for i := 0; i < n; i++ {
		minSum[i] = matrix[0][i]
	}
	for i := 1; i < len(matrix); i++ {
		minSum = min(matrix[i], minSum)
	}
	result := matrix[len(matrix)-1][0]
	for i := 1; i < n; i++ {
		result = min2(result, matrix[len(matrix)-1][i])
	}
	return result
}
