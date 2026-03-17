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
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        if not root:
            return True

        return self._traverse(root.left, root.right)

    def _traverse(self, node_a: Optional[TreeNode], node_b: Optional[TreeNode]) -> bool:
        if not node_a and not node_b:
            return True

        if not node_a or not node_b:
            return False

        if node_a.val != node_b.val:
            return False

        return self._traverse(node_a.left, node_b.right) and self._traverse(node_a.right, node_b.left)
