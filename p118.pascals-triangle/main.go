package main

func generate(numRows int) [][]int {
	res := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
		{1, 5, 10, 10, 5, 1},
		{1, 6, 15, 20, 15, 6, 1},
		{1, 7, 21, 35, 35, 21, 7, 1},
		{1, 8, 28, 56, 70, 56, 28, 8, 1},
		{1, 9, 36, 84, 126, 126, 84, 36, 9, 1},
	}
	if numRows < len(res) {
		return res[:numRows]
	}
	for i := 11; i <= numRows; i++ {
		res = append(res, []int{})
		res[i-1] = make([]int, i)
		res[i-1][0], res[i-1][i-1] = 1, 1
		for j := 1; j <= i-2; j++ {
			res[i-1][j] = res[i-2][j-1] + res[i-2][j]
		}
	}
	return res
}
