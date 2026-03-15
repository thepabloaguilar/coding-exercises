from typing import Optional


class ListNode:
    def __init__(self, val: int = 0, nex: Optional[ListNode] = None) -> None:
        self.val = val
        self.next = nex


class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        result = head = ListNode()

        while list1 or list2:
            num1 = list1.val if list1 else 101
            num2 = list2.val if list2 else 101

            if num1 < num2:
                result.next = ListNode(num1)
                list1 = list1.next
            else:
                result.next = ListNode(num2)
                list2 = list2.next

            result = result.next

        return head.next
