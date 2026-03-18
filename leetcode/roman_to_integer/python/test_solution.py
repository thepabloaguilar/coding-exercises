import pytest

from leetcode.roman_to_integer.python.solution import Solution


@pytest.mark.parametrize(
    ('s', 'expected'),
    (
        ('III', 3),
        ('LVIII', 58),
        ('MCMXCIV', 1994),
    ),
)
def test_base_cases(s: str, expected: int):
    solution = Solution()
    assert solution.romanToInt(s) == expected
