package golang

func Exist(board [][]byte, word string) bool {
	for row := range board {
		for column := range board[row] {
			if dfs(board, row, column, word, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, row int, column int, word string, word_idx int) bool {
	if word_idx == len(word) {
		return true
	}

	if row < 0 || row >= len(board) || column < 0 || column >= len(board[row]) {
		return false
	}

	if board[row][column] != word[word_idx] {
		return false
	}

	letter := board[row][column]
	board[row][column] = '*'

	result := dfs(board, row+1, column, word, word_idx+1) ||
		dfs(board, row-1, column, word, word_idx+1) ||
		dfs(board, row, column+1, word, word_idx+1) ||
		dfs(board, row, column-1, word, word_idx+1)

	board[row][column] = letter

	return result
}
