from typing import Optional

import pytest

from leetcode.same_tree.python.solution import Solution, TreeNode


@pytest.mark.parametrize(
    ('tree_a', 'tree_b', 'expected'),
    [
        (
            TreeNode(
                val=1,
                left=TreeNode(val=2),
                right=TreeNode(val=3),
            ),
            TreeNode(
                val=1,
                left=TreeNode(val=2),
                right=TreeNode(val=3),
            ),
            True,
        ),
        (
            TreeNode(
                val=1,
                left=TreeNode(val=2),
            ),
            TreeNode(
                val=1,
                right=TreeNode(val=2),
            ),
            False,
        ),
        (
            TreeNode(
                val=1,
                left=TreeNode(val=2),
                right=TreeNode(val=1),
            ),
            TreeNode(
                val=1,
                left=TreeNode(val=1),
                right=TreeNode(val=2),
            ),
            False,
        ),
        (None, None, True),
    ]
)
def test_base_cases(tree_a: Optional[TreeNode], tree_b: Optional[TreeNode], expected: bool):
    solution = Solution()
    assert solution.isSameTree(tree_a, tree_b) is expected
