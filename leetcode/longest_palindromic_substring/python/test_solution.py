import pytest

from leetcode.longest_palindromic_substring.python.solution import Solution


@pytest.mark.parametrize(
    ('s', 'expected'),
    (
        ('babad', 'bab'),
        ('cbbd', 'bb'),
    ),
)
def test_base_cases(s: str, expected: str):
    solution = Solution()
    assert solution.longestPalindrome(s) == expected
