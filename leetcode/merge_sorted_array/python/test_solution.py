import pytest

from leetcode.merge_sorted_array.python.solution import Solution


@pytest.mark.parametrize(
    ('nums1', 'm', 'nums2', 'n', 'expected'),
    [
        ([1,2,3,0,0,0], 3, [2,5,6], 3, [1,2,2,3,5,6]),
        ([1], 1, [], 0, [1]),
        ([0], 0, [1], 1, [1]),
    ],
)
def test_base_cases(nums1: list[int], m: int, nums2: list[int], n: int, expected: list[int]) -> None:
    solution = Solution()
    solution.merge(nums1, m, nums2, n)
    assert nums1 == expected
