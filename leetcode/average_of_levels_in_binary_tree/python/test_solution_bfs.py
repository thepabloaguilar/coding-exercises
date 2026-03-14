import pytest

from leetcode.average_of_levels_in_binary_tree.python.solution_bfs import Solution, TreeNode


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
            [3.0,14.5,11.0],
        ),
        (
            TreeNode(
                val=3,
                left=TreeNode(
                    val=9,
                    left=TreeNode(val=15),
                    right=TreeNode(val=7),
                ),
                right=TreeNode(val=20),
            ),
            [3.0,14.5,11.0],
        )
    ]
)
def test_base_cases(data: TreeNode, expected: list[float]):
    solution = Solution()
    assert solution.averageOfLevels(data) == expected
