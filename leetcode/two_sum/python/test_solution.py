import pytest

from leetcode.two_sum.python.solution import Solution


@pytest.mark.parametrize(
    ('nums', 'target', 'expected'),
    [
        ([2,7,11,15], 9, [0, 1]),
        ([3,2,4], 6, [1, 2]),
        ([3,3], 6, [0, 1]),
        ([-3,4,3,90], 0, [0, 2]),
    ]
)
def test_base_cases(nums: list[int], target: int, expected: list[int]):
    solution = Solution()
    assert solution.twoSum(nums, target) == expected
