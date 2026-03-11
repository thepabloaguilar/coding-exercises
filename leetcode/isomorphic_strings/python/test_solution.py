import pytest

from leetcode.isomorphic_strings.python.solution import Solution


@pytest.mark.parametrize(
    ('s', 't', 'expected'),
    (
        ('egg', 'add', True),
        ('f11', 'b23', False),
        ('paper', 'title', True),
        ('badc', 'baba', False),
    ),
)
def test_base_cases(s: str, t: str, expected: bool):
    solution = Solution()
    assert solution.isIsomorphic(s, t) is expected
