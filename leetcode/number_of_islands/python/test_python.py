import pytest

from leetcode.number_of_islands.python.solution import Solution


@pytest.mark.parametrize(
    ('grid', 'expected'),
    [
        (
            [
                ['1','1','1','1','0'],
                ['1','1','0','1','0'],
                ['1','1','0','0','0'],
                ['0','0','0','0','0'],
            ],
            1,
        ),
        (
            [
                ['1','1','0','0','0'],
                ['1','1','0','0','0'],
                ['0','0','1','0','0'],
                ['0','0','0','1','1'],
            ],
            3,
        ),
    ],
)
def test_base_cases(grid: list[list[str]], expected: int) -> None:
    solution = Solution()
    assert solution.numIslands(grid) == expected
