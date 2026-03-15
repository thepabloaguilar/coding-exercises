from typing import Optional


class ListNode:
    def __init__(self, val: int = 0, nex: Optional[ListNode] = None) -> None:
        self.val = val
        self.next = nex


class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        first = l1
        second = l2

        head = result = ListNode()
        rest = 0
        while first or second:
            first_number = first.val if first else 0
            second_number = second.val if second else 0

            summation = first_number + second_number + rest
            value = summation % 10
            rest = summation // 10

            result.next = ListNode(val=value)
            result = result.next

            first = first.next if first else None
            second = second.next if second else None

        if rest > 0:
            result.next = ListNode(val=rest)

        return head.next
