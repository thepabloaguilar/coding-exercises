from typing import Optional

import pytest

from leetcode.minimum_absolute_difference_in_bst.python.solution import Solution, TreeNode


@pytest.mark.parametrize(
    ('root', 'expected'),
    [
        (
            TreeNode(
                val=4,
                left=TreeNode(
                    val=2,
                    left=TreeNode(val=1),
                    right=TreeNode(val=3),
                ),
                right=TreeNode(val=6)
            ),
            1,
        ),
        (
            TreeNode(
                val=1,
                left=TreeNode(val=0),
                right=TreeNode(
                    val=48,
                    left=TreeNode(val=12),
                    right=TreeNode(val=49),
                ),
            ),
            1,
        ),
        (
            TreeNode(
                val=1,
                right=TreeNode(
                    val=5,
                    left=TreeNode(val=3),
                ),
            ),
            2,
        ),
        (
            TreeNode(
                val=236,
                left=TreeNode(
                    val=104,
                    right=TreeNode(val=227),
                ),
                right=TreeNode(
                    val=701,
                    right=TreeNode(val=911),
                ),
            ),
            9,
        ),
    ]
)
def test_base_cases(root: Optional[TreeNode], expected: int):
    solution = Solution()
    assert solution.getMinimumDifference(root) == expected
