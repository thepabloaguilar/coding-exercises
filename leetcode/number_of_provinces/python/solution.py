class Solution:
    def findCircleNum(self, is_connected: list[list[int]]) -> int:
        size = len(is_connected)
        number_of_components = 0
        visited = set()

        for idx in range(size):
            if idx not in visited:
                number_of_components += 1
                self._dfs(idx, is_connected, visited)

        return number_of_components

    def _dfs(
        self,
        node_idx: int,
        is_connected: list[list[int]],
        visited: set[int],
    ) -> None:
        visited.add(node_idx)
        for idx in range(len(is_connected)):
            if is_connected[node_idx][idx] and idx not in visited:
                self._dfs(idx, is_connected, visited)
