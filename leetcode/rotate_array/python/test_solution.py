import pytest

from leetcode.rotate_array.python.solution import Solution


@pytest.mark.parametrize(
    ('nums', 'k', 'expected'),
    [
        ([1,2,3,4,5,6,7], 3, [5,6,7,1,2,3,4]),
        ([-1,-100,3,99], 2, [3,99,-1,-100]),
        ([1, 2], 7, [2,1])
    ]
)
def test_base_cases(nums: list[int], k: int, expected: list[int]):
    solution = Solution()

    solution.rotate(nums, k)

    assert nums == expected
