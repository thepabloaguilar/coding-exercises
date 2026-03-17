import pytest

from leetcode.valid_palindrome.python.solution import Solution


@pytest.mark.parametrize(
    ('s', 'expected'),
    [
        ('A man, a plan, a canal: Panama', True),
        ('race a car', False),
        (' ', True),
        (',,,,,,,,,,,,acva', False),
    ],
)
def test_base_cases(s: str, expected: bool):
    solution = Solution()
    assert solution.isPalindrome(s) is expected
