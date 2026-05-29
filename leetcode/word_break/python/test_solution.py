import pytest

from leetcode.word_break.python.solution import Solution


@pytest.mark.parametrize(
    ('s', 'wordDict', 'expected'),
    [
        ('leetcode', ['leet', 'code'], True),
        ('applepenapple', ['apple', 'pen'], True),
        ('catsandog', ['cats', 'dog', 'sand', 'and', 'cat'], False),
        ('cars', ['car', 'ca', 'rs'], True),
        ('aaaaaaa', ['aaaa', 'aaa'], True),
    ]
)
def test_base_cases(s: str, wordDict: list[str], expected: bool):
    solution = Solution()
    assert solution.wordBreak(s, wordDict) is expected
