class Solution:
    def pacificAtlantic(self, heights: list[list[int]]) -> list[list[int]]:
        if not heights or not heights[0]:
            return []

        row, column = len(heights), len(heights[0])
        pacificReachable = [[False] * column for _ in range(row)]
        atlanticReachable = [[False] * column for _ in range(row)]

        # Check Pacific Ocean connections
        for idx in range(row):
            self._dfs(heights, idx, 0, pacificReachable)
        for idx in range(column):
            self._dfs(heights, 0, idx, pacificReachable)

        # Check Atlantic Ocean connections
        for idx in range(row):
            self._dfs(heights, idx, column - 1, atlanticReachable)
        for idx in range(column):
            self._dfs(heights, row - 1, idx, atlanticReachable)

        # Check overall reachability
        result = []
        for rowIdx in range(row):
            for columnIdx in range(column):
                if pacificReachable[rowIdx][columnIdx] and atlanticReachable[rowIdx][columnIdx]:
                    result.append([rowIdx, columnIdx])

        return result

    def _dfs(self, heights: list[list[int]], row: int, column: int, reachable: list[list[bool]]) -> None:
        reachable[row][column] = True

        directions = [
            (1, 0), # Up
            (-1, 0), # Down
            (0, -1), # Left
            (0, 1), # Right
        ]
        for rowDirection, columnDirection in directions:
            newRow, newColumn = row + rowDirection, column + columnDirection

            isWithinBoundaries = 0 <= newRow < len(heights) and 0 <= newColumn < len(heights[0])
            if isWithinBoundaries and not reachable[newRow][newColumn]:
                if heights[newRow][newColumn] >= heights[row][column]:
                    self._dfs(heights, newRow, newColumn, reachable)
