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
    def getMinimumDifference(self, root: Optional[TreeNode]) -> int:
        return self._traverse(root, float('-inf'), float('inf'))

    def _traverse(self, node: Optional[TreeNode], low: int | float, high: int | float) -> int:
        if not node:
            return high - low

        left = self._traverse(node.left, low, node.val)
        right = self._traverse(node.right, node.val, high)

        return min(left, right)
