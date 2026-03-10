import pytest

from leetcode.majority_element.python.solution import Solution


@pytest.mark.parametrize(
    ('nums', 'expected'),
    [
        ([3,2,3], 3),
        ([2,2,1,1,1,2,2], 2),
        ([10,9,9,9,10], 9),
    ]
)
def test_base_cases(nums: list[int], expected: int):
    solution = Solution()
    assert solution.majorityElement(nums) == expected
