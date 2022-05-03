package main

func dfs(board [][]byte, filter []bool, i, j int, wordBytes []byte) bool {
	if len(wordBytes) == 0 || (len(wordBytes) == 1 && wordBytes[0] == board[i][j]) {
		return true
	} else if wordBytes[0] != board[i][j] {
		return false
	}
	ny := len(board[0])

	candidates := [4][2]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	for k := range candidates {
		if candidates[k][0] < 0 || candidates[k][0] >= len(board) || candidates[k][1] < 0 || candidates[k][1] >= len(board[0]) || filter[candidates[k][0]*ny+candidates[k][1]] {
			continue
		}
		filter[candidates[k][0]*ny+candidates[k][1]] = true
		if dfs(board, filter, candidates[k][0], candidates[k][1], wordBytes[1:]) {
			return true
		}
		filter[candidates[k][0]*ny+candidates[k][1]] = false
	}
	return false
}

func exist(board [][]byte, word string) bool {
	wordBytes := []byte(word)
	nx := len(board)
	ny := len(board[0])
	for i := 0; i < nx; i++ {
		for j := 0; j < ny; j++ {
			if wordBytes[0] != board[i][j] {
				continue
			}
			filter := make([]bool, nx*ny)
			filter[i*ny+j] = true
			if dfs(board, filter, i, j, wordBytes) {
				return true
			}
		}
	}
	return false
}
