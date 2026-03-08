class Solution:
    def numEnclaves(self, grid: list[list[int]]) -> int:
        if not grid:
            return 0

        not_connected = len(grid) * len(grid[0])
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                item = grid[i][j]
                if item == 0:
                    not_connected -= 1

                is_boundary = i - 1 < 0 or i == len(grid) - 1
                is_boundary = is_boundary or j - 1 < 0 or j == len(grid[0]) - 1
                if is_boundary and item == 1:
                    not_connected -= self._check(grid, i, j)

        return not_connected

    def _check(self, grid: list[list[int]], i: int, j: int) -> int:
        is_out_of_boundry = i < 0 or i >= len(grid) \
                            or j < 0 or j >= len(grid[0])
        if is_out_of_boundry or grid[i][j] != 1:
            return 0

        grid[i][j] = 2

        connected = 1
        connected += self._check(grid, i + 1, j)
        connected += self._check(grid, i - 1, j)
        connected += self._check(grid, i, j + 1)
        connected += self._check(grid, i, j - 1)

        return connected
