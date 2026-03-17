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
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        return self._dfs(root)

    def _dfs(self, node: Optional[TreeNode], level: int = 0) -> int:
        if not node:
            return level

        return max(
            self._dfs(node.left, level + 1),
            self._dfs(node.right, level + 1),
        )
