from typing import Optional

import pytest

from leetcode.count_complete_tree_nodes.python.solution import Solution, TreeNode


@pytest.mark.parametrize(
    ('root', 'expected'),
    [
        (
            TreeNode(
                val=1,
                left=TreeNode(
                    val=2,
                    left=TreeNode(val=4),
                    right=TreeNode(val=5),
                ),
                right=TreeNode(
                    val=3,
                    left=TreeNode(val=6),
                ),
            ),
            6,
        ),
        (None, 0),
        (TreeNode(val=1), 1),
    ]
)
def test_base_cases(root: Optional[TreeNode], expected: int):
    solution = Solution()
    assert solution.countNodes(root) == expected
