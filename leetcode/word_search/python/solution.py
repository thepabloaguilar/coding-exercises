class Solution:
    def exist(self, board: list[list[str]], word: str) -> bool:
        for row in range(len(board)):
            for column in range(len(board[row])):
                if self._dfs(board, row, column, word, 0):
                    return True

        return False

    def _dfs(self, board: list[list[str]], row: int, column: int, word: str, word_idx: int):
        if row < 0 or row >= len(board) or column < 0 or column >= len(board[row]):
            return False

        if board[row][column] != word[word_idx]:
            return False

        if word_idx >= len(word) - 1:
            return True

        letter = board[row][column]
        board[row][column] = '*'

        result = (
                self._dfs(board, row + 1, column, word, word_idx + 1)
                or self._dfs(board, row - 1, column, word, word_idx + 1)
                or self._dfs(board, row, column + 1, word, word_idx + 1)
                or self._dfs(board, row, column - 1, word, word_idx + 1)
        )

        board[row][column] = letter

        return result
