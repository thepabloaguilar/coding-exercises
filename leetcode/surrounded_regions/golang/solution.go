package golang

func Solve(board [][]byte) {
	for i := range board {
		if board[i][0] == 'O' {
			capture(board, i, 0)
		}
		if board[i][len(board[0])-1] == 'O' {
			capture(board, i, len(board[0])-1)
		}
	}

	for i := range board[0] {
		if board[0][i] == 'O' {
			capture(board, 0, i)
		}
		if board[len(board)-1][i] == 'O' {
			capture(board, len(board)-1, i)
		}
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'P' {
				board[i][j] = 'O'
			}
		}
	}
}

func capture(board [][]byte, i int, j int) {
	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) || board[i][j] == 'X' || board[i][j] == 'P' {
		return
	}

	board[i][j] = 'P'

	capture(board, i+1, j)
	capture(board, i-1, j)
	capture(board, i, j+1)
	capture(board, i, j-1)
}
