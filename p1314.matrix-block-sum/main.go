package main

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func sum(mat [][]int, i, j, k int) int {
	il := max(0, i-k)
	ir := min(len(mat)-1, i+k)
	jl := max(0, j-k)
	jr := min(len(mat[0])-1, j+k)

	total := 0

	for i := il; i <= ir; i++ {
		for j := jl; j <= jr; j++ {
			total += mat[i][j]
		}

	}
	return total

}
func matrixBlockSum(mat [][]int, k int) [][]int {
	res := make([][]int, len(mat))

	for i := 0; i < len(mat); i++ {
		res[i] = make([]int, len(mat[i]))
		for j := 0; j < len(mat[i]); j++ {
			res[i][j] = sum(mat, i, j, k)
		}
	}
	return res
}
