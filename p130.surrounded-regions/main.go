package main

func dfs(board [][]byte, x, y int) {
	nx := len(board)
	ny := len(board[0])

	board[x][y] = 'K'
	if x > 0 && board[x-1][y] == 'O' {
		dfs(board, x-1, y)
	}
	if x < nx-1 && board[x+1][y] == 'O' {
		dfs(board, x+1, y)
	}
	if y > 0 && board[x][y-1] == 'O' {
		dfs(board, x, y-1)
	}
	if y < ny-1 && board[x][y+1] == 'O' {
		dfs(board, x, y+1)
	}
}

func solve(board [][]byte) {
	nx := len(board)
	ny := len(board[0])

	for i := 0; i < nx; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][ny-1] == 'O' {
			dfs(board, i, ny-1)
		}
	}
	for j := 0; j < ny; j++ {
		if board[0][j] == 'O' {
			dfs(board, 0, j)
		}
		if board[nx-1][j] == 'O' {
			dfs(board, nx-1, j)
		}
	}
	for i := 0; i < nx; i++ {
		for j := 0; j < ny; j++ {
			if board[i][j] == 'K' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}
