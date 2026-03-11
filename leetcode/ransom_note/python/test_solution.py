import pytest

from leetcode.ransom_note.python.solution import Solution


@pytest.mark.parametrize(
    ('ransomNote', 'magazine', 'expected'),
    (
        ('a', 'b', False),
        ('aa', 'ab', False),
        ('aa', 'aab', True),
    ),
)
def test_base_cases(ransomNote: str, magazine: str, expected: bool):
    solution = Solution()
    assert solution.canConstruct(ransomNote, magazine) is expected
