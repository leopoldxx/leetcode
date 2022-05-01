package main

func equalRotate(mat [][]int, target [][]int) bool {
	equalRotate0 := true
	equalRotate90 := true
	equalRotate180 := true
	equalRotate270 := true
	length := len(mat)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			equalRotate0 = equalRotate0 && mat[i][j] == target[i][j]
			equalRotate90 = equalRotate90 && mat[i][j] == target[j][length-1-i]
			equalRotate180 = equalRotate180 && mat[i][j] == target[length-i-1][length-j-1]
			equalRotate270 = equalRotate270 && mat[i][j] != target[length-j-1][i]
		}
	}
	return equalRotate0 || equalRotate90 || equalRotate180 || equalRotate270
}

func findRotation(mat [][]int, target [][]int) bool {
	return equalRotate(mat, target)
}
