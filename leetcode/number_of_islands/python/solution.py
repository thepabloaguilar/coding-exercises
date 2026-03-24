class Solution:
    def numIslands(self, grid: list[list[str]]) -> int:
        islands = 0

        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == '1':
                    self._traverse(grid, i, j)
                    islands += 1

        return islands

    def _traverse(self, grid: list[list[str]], i: int, j: int) -> None:
        if (i < 0 or i >= len(grid)) or (j < 0 or j >= len(grid[0])) or grid[i][j] == '0':
            return

        grid[i][j] = '0'

        self._traverse(grid, i + 1, j)
        self._traverse(grid, i - 1, j)
        self._traverse(grid, i, j + 1)
        self._traverse(grid, i, j - 1)
