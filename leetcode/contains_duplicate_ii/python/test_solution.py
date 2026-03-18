import pytest

from leetcode.contains_duplicate_ii.python.solution import Solution


@pytest.mark.parametrize(
    ('nums', 'k', 'expected'),
    [
        ([1,2,3,1], 3, True),
        ([1,0,1,1], 1, True),
        ([1,2,3,1,2,3], 2, False),
        ([1,2,3,1,2,3,2], 2, True),
    ]
)
def test_base_cases(nums: list[int], k: int, expected: bool):
    solution = Solution()
    assert solution.containsNearbyDuplicate(nums, k) is expected
