import pytest

from leetcode.subarray_sum_equals_k.python.solution import Solution


@pytest.mark.parametrize(
    ('nums', 'k', 'expected'),
    (
        ([1,1,1], 2, 2),
        ([1,2,3], 3, 2),
    ),
)
def test_base_cases(nums: list[int], k: int, expected: int):
    solution = Solution()
    assert solution.subarraySum(nums, k) == expected
