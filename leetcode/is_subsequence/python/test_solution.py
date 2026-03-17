import pytest

from leetcode.is_subsequence.python.solution import Solution


@pytest.mark.parametrize(
    ('s', 't', 'expected'),
    [
        ('abc', 'ahbgdc', True),
        ('axc', 'ahbgdc', False),
        ('', 'ahbgdc', True),
    ],
)
def test_base_cases(s: str, t: str, expected: bool):
    solution = Solution()
    assert solution.isSubsequence(s, t) is expected
