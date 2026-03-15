from collections import deque
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
    def levelOrderBottom(self, root: Optional[TreeNode]) -> list[list[int]]:
        if not root:
            return []

        queue = deque([root])
        values = []
        while len(queue) > 0:
            level_values = []
            for _ in range(len(queue)):
                node = queue.popleft()
                level_values.append(node.val)

                if node.left: queue.append(node.left)
                if node.right: queue.append(node.right)

            values.append(level_values)

        return values[::-1]
