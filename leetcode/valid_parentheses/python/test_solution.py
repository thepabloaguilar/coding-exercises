import pytest

from leetcode.valid_parentheses.python.solution import Solution


@pytest.mark.parametrize(
    ('s', 'expected'),
    (
        ('()', True),
        ('()[]{}', True),
        ('(]', False),
        ('([])', True),
        ('([)]', False),
    ),
)
def test_base_cases(s: str, expected: bool):
    solution = Solution()
    assert solution.isValid(s) is expected
