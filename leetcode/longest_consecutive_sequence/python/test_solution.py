import pytest

from leetcode.longest_consecutive_sequence.python.solution import Solution


@pytest.mark.parametrize(
    ('nums', 'expected'),
    (
        ([100,4,200,1,3,2], 4),
        ([0,3,7,2,5,8,4,6,0,1], 9),
        ([1,0,1,2], 3),
        ([0], 1),
        ([0,-1], 2),
        ([], 0),
    ),
)
def test_base_cases(nums: list[int], expected: int):
    solution = Solution()
    assert solution.longestConsecutive(nums) == expected
