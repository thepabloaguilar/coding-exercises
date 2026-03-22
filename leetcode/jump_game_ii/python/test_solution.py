import pytest

from leetcode.jump_game_ii.python.solution import Solution


@pytest.mark.parametrize(
    ('nums', 'expected'),
    (
        ([2,3,1,1,4], 2),
        ([2,3,0,1,4], 2),
        ([2,1], 1),
        ([1,2,3], 2),
        ([3,2,1], 1),
    ),
)
def test_base_cases(nums: list[int], expected: int):
    solution = Solution()
    assert solution.jump(nums) == expected
