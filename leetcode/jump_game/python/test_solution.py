import pytest

from leetcode.jump_game.python.solution import Solution


@pytest.mark.parametrize(
    ('nums', 'expected'),
    (
        ([2,3,1,1,4], True),
        ([3,2,1,0,4], False),
        ([0], True),
        ([2,1], True),
        ([2,4,1,1,4], True),
        ([1,1,1,0], True),
    ),
)
def test_base_cases(nums: list[int], expected: bool):
    solution = Solution()
    assert solution.canJump(nums) == expected
