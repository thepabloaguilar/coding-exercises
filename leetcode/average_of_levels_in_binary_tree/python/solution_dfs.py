from typing import Optional


class TreeNode:
    def __init__(
        self,
        val: int,
        left: Optional[TreeNode] = None,
        right: Optional[TreeNode] = None,
    ) -> None:
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def averageOfLevels(self, root: Optional[TreeNode]) -> list[float]:
        numbers_by_level = []
        self._dfs(root, 0, numbers_by_level)

        return [
            sum(numbers) / len(numbers) for numbers in numbers_by_level
        ]

    def _dfs(self, root: Optional[TreeNode], level: int, numbers_by_level: list[list[int]]) -> None:
        if not root:
            return

        if level == len(numbers_by_level):
            numbers_by_level.append([])

        numbers_by_level[level].append(root.val)

        self._dfs(root.left, level + 1, numbers_by_level)
        self._dfs(root.right, level + 1, numbers_by_level)
