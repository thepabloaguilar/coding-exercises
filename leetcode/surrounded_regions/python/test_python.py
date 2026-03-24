import pytest

from leetcode.surrounded_regions.python.solution import Solution


@pytest.mark.parametrize(
    ('board', 'expected'),
    [
        (
            [
                ['X','X','X','X'],
                ['X','O','O','X'],
                ['X','X','O','X'],
                ['X','O','X','X'],
            ],
            [
                ['X','X','X','X'],
                ['X','X','X','X'],
                ['X','X','X','X'],
                ['X','O','X','X'],
            ],
        ),
        (
            [['X']], [['X']],
        ),
        (
            [
                ['O','O','O'],
                ['O','O','O'],
                ['O','O','O'],
            ],
            [
                ['O','O','O'],
                ['O','O','O'],
                ['O','O','O'],
            ],
        ),
        (
            [
                ['O','O','O','O','X','X'],
                ['O','O','O','O','O','O'],
                ['O','X','O','X','O','O'],
                ['O','X','O','O','X','O'],
                ['O','X','O','X','O','O'],
                ['O','X','O','O','O','O'],
            ],
            [
                ['O','O','O','O','X','X'],
                ['O','O','O','O','O','O'],
                ['O','X','O','X','O','O'],
                ['O','X','O','O','X','O'],
                ['O','X','O','X','O','O'],
                ['O','X','O','O','O','O'],
            ],
        ),
        (
            [
                ['X','O','X'],
                ['O','X','O'],
                ['X','O','X'],
            ],
            [
                ['X','O','X'],
                ['O','X','O'],
                ['X','O','X'],
            ],
        ),
        (
            [
                ['X','X','X','X','O'],
                ['X','O','O','X','X'],
                ['O','X','X','X','X'],
            ],
            [
                ['X','X','X','X','O'],
                ['X','X','X','X','X'],
                ['O','X','X','X','X'],
            ],
        ),
    ],
)
def test_base_cases(board: list[list[str]], expected: list[list[str]]) -> None:
    solution = Solution()

    solution.solve(board)

    assert board == expected
