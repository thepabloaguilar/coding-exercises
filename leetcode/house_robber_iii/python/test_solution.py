from typing import Optional

import pytest

from leetcode.house_robber_iii.python.solution import Solution, TreeNode


@pytest.mark.parametrize(
    ('root', 'expected'),
    [
        (
            TreeNode(
                val=3,
                left=TreeNode(
                    val=2,
                    right=TreeNode(val=3),
                ),
                right=TreeNode(
                    val=3,
                    right=TreeNode(val=1),
                ),
            ),
            7,
        ),
        (
            TreeNode(
                val=3,
                left=TreeNode(
                    val=4,
                    left=TreeNode(val=1),
                    right=TreeNode(val=3),
                ),
                right=TreeNode(
                    val=5,
                    right=TreeNode(val=1),
                ),
            ),
            9
        )
    ]
)
def test_base_cases(root: Optional[TreeNode], expected: int):
    solution = Solution()
    assert solution.rob(root) == expected
