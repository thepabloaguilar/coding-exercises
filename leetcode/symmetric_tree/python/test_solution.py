from typing import Optional

import pytest

from leetcode.symmetric_tree.python.solution import Solution, TreeNode


@pytest.mark.parametrize(
    ('tree', 'expected'),
    [
        (
            TreeNode(
                val=1,
                left=TreeNode(
                    val=2,
                    left=TreeNode(val=3),
                    right=TreeNode(val=4)
                ),
                right=TreeNode(
                    val=2,
                    left=TreeNode(val=4),
                    right=TreeNode(val=3)
                )
            ),
            True,
        ),
        (
            TreeNode(
                val=1,
                left=TreeNode(
                    val=2,
                    right=TreeNode(val=3),
                ),
                right=TreeNode(
                    val=2,
                    right=TreeNode(val=3),
                ),
            ),
            False
        ),
        (None, True),
    ]
)
def test_base_cases(tree: Optional[TreeNode], expected: bool):
    solution = Solution()
    assert solution.isSymmetric(tree) is expected
