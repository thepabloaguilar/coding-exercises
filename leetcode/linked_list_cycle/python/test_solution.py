from leetcode.linked_list_cycle.python.solution import Solution, ListNode


def test_base_cases_first():
    head = ListNode(3)
    head.next = ListNode(2)
    head.next.next = ListNode(0)
    head.next.next.next = ListNode(-4)

    # Link node 1 to node 3
    head.next.next.next.next = head.next

    solution = Solution()
    assert solution.hasCycle(head) is True


def test_base_cases_second():
    head = ListNode(1)
    head.next = ListNode(2)

    # Link node 0 to node 1
    head.next.next = head

    solution = Solution()
    assert solution.hasCycle(head) is True


def test_base_cases_third():
    head = ListNode(1)

    solution = Solution()
    assert solution.hasCycle(head) is False
