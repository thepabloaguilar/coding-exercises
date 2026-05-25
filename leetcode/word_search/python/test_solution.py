import pytest

from leetcode.word_search.python.solution import Solution


@pytest.mark.parametrize(
    ('board', 'word', 'expected'),
    [
        (
            [
                ['A', 'B', 'C', 'E'],
                ['S', 'F', 'C', 'S'],
                ['A', 'D', 'E', 'E'],
            ],
            'ABCCED',
            True,
        ),
        (
            [
                ['A', 'B', 'C', 'E'],
                ['S', 'F', 'C', 'S'],
                ['A', 'D', 'E', 'E'],
            ],
            'SEE',
            True,
        ),
        (
            [
                ['A', 'B', 'C', 'E'],
                ['S', 'F', 'C', 'S'],
                ['A', 'D', 'E', 'E'],
            ],
            'ABCB',
            False,
        ),
        ([['a']], 'a', True),
    ]
)
def test_base_cases(board: list[list[str]], word: str, expected: bool):
    solution = Solution()
    assert solution.exist(board, word) is expected
