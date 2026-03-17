from typing import Optional

import pytest

from leetcode.maximum_depth_of_binary_tree.python.solution import Solution, TreeNode


@pytest.mark.parametrize(
    ('data', 'expected'),
    [
        (
            TreeNode(
                val=3,
                left=TreeNode(val=9),
                right=TreeNode(
                    val=20,
                    left=TreeNode(val=15),
                    right=TreeNode(val=7),
                ),
            ),
            3,
        ),
        (
            TreeNode(
                val=1,
                right=TreeNode(val=2),
            ),
            2,
        ),
        (None, 0),
    ]
)
def test_base_cases(data: Optional[TreeNode], expected: int):
    solution = Solution()
    assert solution.maxDepth(data) == expected
