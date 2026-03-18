import pytest

from leetcode.happy_number.python.solution import Solution


@pytest.mark.parametrize(
    ('n', 'expected'),
    (
        (19, True),
        (2, False),
    ),
)
def test_base_cases(n: int, expected: bool):
    solution = Solution()
    assert solution.isHappy(n) is expected
