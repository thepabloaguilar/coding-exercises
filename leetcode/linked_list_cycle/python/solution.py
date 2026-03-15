from typing import Optional


class ListNode:
    def __init__(self, val: int) -> None:
        self.val = val
        self.next = None


class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        node = head
        seen = set()
        while node:
            if node in seen:
                return True

            seen.add(node)
            node = node.next

        return False
