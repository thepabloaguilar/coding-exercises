from typing import Optional

import pytest

from leetcode.binary_tree_right_side_view.python.solution import Solution, TreeNode


@pytest.mark.parametrize(
    ('data', 'expected'),
    [
        (
            TreeNode(
                val=1,
                left=TreeNode(
                    val=2,
                    right=TreeNode(val=5),
                ),
                right=TreeNode(
                    val=3,
                    right=TreeNode(val=4),
                ),
            ),
            [1,3,4],
        ),
        (
            TreeNode(
                val=1,
                left=TreeNode(
                    val=2,
                    left=TreeNode(
                        val=4,
                        left=TreeNode(val=5),
                    ),
                ),
                right=TreeNode(val=3),
            ),
            [1,3,4,5],
        ),
        (
            TreeNode(
                val=1,
                right=TreeNode(val=3),
            ),
            [1,3],
        ),
        (None, []),
    ]
)
def test_base_cases(data: Optional[TreeNode], expected: list[int]):
    solution = Solution()
    assert solution.rightSideView(data) == expected
