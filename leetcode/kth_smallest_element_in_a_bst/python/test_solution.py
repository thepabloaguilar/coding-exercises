from typing import Optional

import pytest

from leetcode.kth_smallest_element_in_a_bst.python.solution import Solution, TreeNode


@pytest.mark.parametrize(
    ('root', 'k', 'expected'),
    [
        (
            TreeNode(
                val=3,
                left=TreeNode(
                    val=1,
                    right=TreeNode(val=2),
                ),
                right=TreeNode(val=4),
            ),
            1,
            1,
        ),
        (
            TreeNode(
                val=5,
                left=TreeNode(
                    val=3,
                    left=TreeNode(
                        val=2,
                        left=TreeNode(val=1),
                    ),
                    right=TreeNode(val=4),
                ),
                right=TreeNode(val=6),
            ),
            3,
            3,
        ),
        (
            TreeNode(
                val=5,
                left=TreeNode(
                    val=3,
                    left=TreeNode(
                        val=2,
                        left=TreeNode(val=1),
                    ),
                    right=TreeNode(val=4),
                ),
                right=TreeNode(val=6),
            ),
            4,
            4,
        ),
        (
            TreeNode(val=5), 1, 5,
        ),
    ]
)
def test_base_cases(root: Optional[TreeNode], k: int, expected: list[int]):
    solution = Solution()
    assert solution.kthSmallest(root, k) == expected
