from typing import Optional

import pytest

from leetcode.binary_tree_level_order_traversal_ii.python.solution import Solution, TreeNode


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
                [[15,7],[9,20],[3]],
        ),
        (
                TreeNode(val=1),
                [[1]],
        ),
        (None, []),
    ]
)
def test_base_cases(data: Optional[TreeNode], expected: list[int]):
    solution = Solution()
    assert solution.levelOrderBottom(data) == expected
