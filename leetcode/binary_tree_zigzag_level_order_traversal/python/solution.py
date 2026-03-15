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
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> list[list[int]]:
        if not root:
            return []

        values = []
        queue = deque([root])
        to_left = True
        while len(queue) > 0:
            level_values = []
            for _ in range(len(queue)):
                node = queue.popleft()

                if to_left:
                    level_values.append(node.val)
                else:
                    level_values.insert(0, node.val)

                if node.left: queue.append(node.left)
                if node.right: queue.append(node.right)

            to_left = not to_left
            values.append(level_values)

        return values
