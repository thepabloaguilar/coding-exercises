import pytest

from leetcode.number_of_provinces.python.solution import Solution


@pytest.mark.parametrize(
    ('data', 'expected'),
    [
        (
            [[1,1,0],[1,1,0],[0,0,1]],
            2,
        ),
        (
            [[1,0,0],[0,1,0],[0,0,1]],
            3,
        ),
    ],
)
def test_base_cases(data: list[list[int]], expected: int) -> None:
    solution = Solution()
    assert solution.findCircleNum(data) == expected
