import pytest

from leetcode.find_the_index_of_the_first_occurrence_in_a_string.python.solution import Solution


@pytest.mark.parametrize(
    ('haystack', 'needle', 'expected'),
    (
        ('sadbutsad', 'sad', 0),
        ('leetcode', 'leeto', -1),
        ('mississippi', 'issip', 4),
        ('a', 'a', 0),
        ('aaa', 'aaaa', -1),
    ),
)
def test_base_cases(haystack: str, needle: str, expected: int):
    solution = Solution()
    assert solution.strStr(haystack, needle) == expected
