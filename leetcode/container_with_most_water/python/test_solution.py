import pytest

from leetcode.container_with_most_water.python.solution import Solution


@pytest.mark.parametrize(
    ('height', 'expected'),
    (
        ([1,8,6,2,5,4,8,3,7], 49),
        ([1,1], 1),
    ),
)
def test_base_cases(height: list[int], expected: int):
    solution = Solution()
    assert solution.maxArea(height) == expected
