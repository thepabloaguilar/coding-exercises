from typing import Optional

import pytest

from leetcode.invert_binary_tree.python.solution import Solution, TreeNode


@pytest.mark.parametrize(
    ('tree', 'expected'),
    [
        (
            TreeNode(
                val=4,
                left=TreeNode(
                    val=2,
                    left=TreeNode(val=1),
                    right=TreeNode(val=3)
                ),
                right=TreeNode(
                    val=7,
                    left=TreeNode(val=6),
                    right=TreeNode(val=9)
                )
            ),
            TreeNode(
                val=4,
                left=TreeNode(
                    val=7,
                    left=TreeNode(val=9),
                    right=TreeNode(val=6)
                ),
                right=TreeNode(
                    val=2,
                    left=TreeNode(val=3),
                    right=TreeNode(val=1)
                )
            ),
        ),
        (
            TreeNode(
                val=2,
                left=TreeNode(val=1),
                right=TreeNode(val=3),
            ),
            TreeNode(
                val=2,
                left=TreeNode(val=3),
                right=TreeNode(val=1),
            ),
        ),
        (None, None),
    ]
)
def test_base_cases(tree: Optional[TreeNode], expected: Optional[TreeNode]):
    solution = Solution()
    assert _check_same_tree(solution.invertTree(tree), expected) is True


def _check_same_tree(node_a: Optional[TreeNode], node_b: Optional[TreeNode]) -> bool:
    if not node_a and not node_b:
        return True

    if not node_a or not node_b:
        return False

    if node_a.val != node_b.val:
        return False

    return _check_same_tree(node_a.left, node_b.left) and _check_same_tree(node_a.right, node_b.right)
