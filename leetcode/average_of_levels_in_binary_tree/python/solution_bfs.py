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
    def averageOfLevels(self, root: Optional[TreeNode]) -> list[float]:
        average_by_level = []

        nodes = deque([root])
        while nodes:
            summation = 0
            size = len(nodes)

            for _ in range(size):
                node = nodes.popleft()
                summation += node.val
                if node.left:
                    nodes.append(node.left)
                if node.right:
                    nodes.append(node.right)

            average_by_level.append(summation / size)

        return average_by_level
