import pytest

from leetcode.trapping_rain_water.python.solution import Solution


@pytest.mark.parametrize(
    ('height', 'expected'),
    (
        ([0,1,0,2,1,0,1,3,2,1,2,1], 6),
        ([4,2,0,3,2,5], 9),
    ),
)
def test_base_cases(height: list[int], expected: int):
    solution = Solution()
    assert solution.trap(height) == expected
