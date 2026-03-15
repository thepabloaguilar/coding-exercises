from typing import Optional

import pytest

from leetcode.merge_two_sorted_lists.python.solution import Solution, ListNode


@pytest.mark.parametrize(
    ('l1', 'l2', 'expected'),
    [
        (
            ListNode(1, ListNode(2, ListNode(4))),
            ListNode(1, ListNode(3, ListNode(4))),
            ListNode(1, ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(4)))))),
        ),
        (None, None, None),
        (None, ListNode(0), ListNode(0))
    ]
)
def test_base_cases(l1: Optional[ListNode], l2: Optional[ListNode], expected: Optional[ListNode]):
    solution = Solution()
    actual = solution.mergeTwoLists(l1, l2)

    assert _list_node_to_list(actual) == _list_node_to_list(expected)


def _list_node_to_list(node: Optional[ListNode]) -> Optional[list[int]]:
    if node is None:
        return None
    return [node.val] + _list_node_to_list(node.next) if node.next else [node.val]
