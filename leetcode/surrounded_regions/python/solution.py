class Solution:
    def solve(self, board: list[list[str]]) -> None:
        # Check left and right edges
        for i in range(len(board)):
            if board[i][0] == 'O':
                self._capture(board, i, 0)
            if board[i][len(board[0])-1] == 'O':
                self._capture(board, i, len(board[0])-1)

        # Check top and bottom edges
        for i in range(len(board[0])):
            if board[0][i] == 'O':
                self._capture(board, 0, i)
            if board[len(board)-1][i] == 'O':
                self._capture(board, len(board)-1, i)

        for i in range(len(board)):
            for j in range(len(board[0])):
                if board[i][j] == 'O':
                    board[i][j] = 'X'
                elif board[i][j] == 'P':
                    board[i][j] = 'O'

    def _capture(self, board: list[list[str]], i: int, j: int) -> None:
        if i < 0 or i == len(board) or j < 0 or j == len(board[0]) or board[i][j] == 'X' or board[i][j] == 'P':
            return

        board[i][j] = 'P'

        self._capture(board, i + 1, j)
        self._capture(board, i - 1, j)
        self._capture(board, i, j + 1)
        self._capture(board, i, j - 1)
