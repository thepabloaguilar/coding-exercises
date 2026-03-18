import pytest

from leetcode.length_of_last_word.python.solution import Solution


@pytest.mark.parametrize(
    ('s', 'expected'),
    (
        ('Hello World', 5),
        ('   fly me   to   the moon  ', 4),
        ('luffy is still joyboy', 6),
        ('Hello', 5),
        ('a ', 1),
    ),
)
def test_base_cases(s: str, expected: int):
    solution = Solution()
    assert solution.lengthOfLastWord(s) == expected
