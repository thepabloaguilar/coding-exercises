import pytest

from leetcode.remove_element.python.solution import Solution


@pytest.mark.parametrize(
    ('nums', 'val', 'expected'),
    [
        ([3,2,2,3], 3, [2,2]),
        ([0,1,2,2,3,0,4,2], 2, [0,1,4,0,3]),
    ]
)
def test_base_cases(nums: list[int], val: int, expected: list[int]):
    solution = Solution()

    remaining_nums = solution.removeElement(nums, val)

    assert remaining_nums == len(expected)
    assert sorted(nums[:remaining_nums]) == sorted(expected)
