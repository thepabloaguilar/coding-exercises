import pytest

from leetcode.pacific_atlantic_water_flow.python.solution import Solution


@pytest.mark.parametrize(
    ('heights', 'expected'),
    [
        (
            [
                [1, 2, 2, 3, 5],
                [3, 2, 3, 4, 4],
                [2, 4, 5, 3, 1],
                [6, 7, 1, 4, 5],
                [5, 1, 1, 2, 4],
            ],
            [[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]],
        ),
        ([[1]], [[0, 0]]),
    ]
)
def test_base_cases(heights: list[list[int]], expected: list[int]):
    solution = Solution()
    assert solution.pacificAtlantic(heights) == expected
