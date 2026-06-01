import pytest

from leetcode.median_of_two_sorted_arrays.python.solution import Solution


@pytest.mark.parametrize(
    ('nums1', 'nums2', 'expected'),
    [
        ([1, 3], [2], 2.0),
        ([1, 2], [3, 4], 2.5),
        ([1, 3, 8, 9, 15], [7, 11, 18, 19, 21, 25], 11.0)
    ]
)
def test_base_cases(nums1: list[int], nums2: list[int], expected: float):
    solution = Solution()
    assert solution.findMedianSortedArrays(nums1, nums2) == expected
