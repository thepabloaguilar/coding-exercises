from typing import Optional

import pytest

from leetcode.path_sum.python.solution import Solution, TreeNode


@pytest.mark.parametrize(
    ('root', 'target_sum', 'expected'),
    [
        (
            TreeNode(
                val=5,
                left=TreeNode(
                    val=4,
                    left=TreeNode(
                        val=11,
                        left=TreeNode(val=7),
                        right=TreeNode(val=2),
                    )
                ),
                right=TreeNode(
                    val=8,
                    left=TreeNode(val=13),
                    right=TreeNode(
                        val=4,
                        right=TreeNode(val=1),
                    ),
                ),
            ),
            22,
            True,
        ),
        (
            TreeNode(
                val=1,
                left=TreeNode(val=2),
                right=TreeNode(val=3),
            ),
            5,
            False,
        ),
        (None, 0, False),
        (
            TreeNode(val=1, left=TreeNode(val=2)),
            1,
            False,
        ),
    ]
)
def test_base_cases(root: Optional[TreeNode], target_sum: int, expected: bool):
    solution = Solution()
    assert solution.hasPathSum(root, target_sum) is expected
