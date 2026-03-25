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
    def rob(self, root: Optional[TreeNode]) -> int:
        return max(self._traverse(root))

    def _traverse(self, node: Optional[TreeNode]) -> tuple[int, int]:
        if not node:
            return 0, 0

        left_left, left_right = self._traverse(node.left)
        right_left, right_right = self._traverse(node.right)

        return (
            max(left_left, left_right) + max(right_left, right_right),
            node.val + left_left + right_left,
        )
