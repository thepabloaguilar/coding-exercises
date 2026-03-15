import pytest

from leetcode.add_two_numbers.python.solution import Solution, ListNode


@pytest.mark.parametrize(
    ('l1', 'l2', 'expected'),
    [
        (
            ListNode(2, ListNode(4, ListNode(3))),
            ListNode(5, ListNode(6, ListNode(4))),
            ListNode(7, ListNode(0, ListNode(8))),
        ),
        (
            ListNode(0), ListNode(0), ListNode(0),
        ),
        (
            ListNode(9, ListNode(9, ListNode(9, ListNode(9, ListNode(9, ListNode(9, ListNode(9))))))),
            ListNode(9, ListNode(9, ListNode(9, ListNode(9)))),
            ListNode(8, ListNode(9, ListNode(9, ListNode(9, ListNode(0, ListNode(0, ListNode(0, ListNode(1)))))))),
        )
    ]
)
def test_base_cases(l1: ListNode, l2: ListNode, expected: ListNode):
    solution = Solution()
    actual = solution.addTwoNumbers(l1, l2)

    assert _list_node_to_list(actual) == _list_node_to_list(expected)


def _list_node_to_list(node: ListNode) -> list[int]:
    return [node.val] + _list_node_to_list(node.next) if node.next else [node.val]
