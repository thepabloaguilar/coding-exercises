import pytest

from leetcode.candy.python.solution import Solution


@pytest.mark.parametrize(
    ('ratings', 'expected'),
    (
        ([1,0,2], 5),
        ([1,2,2], 4),
        ([1,2,87,87,87,2,1], 13),
    ),
)
def test_base_cases(ratings: list[int], expected: int):
    solution = Solution()
    assert solution.candy(ratings) == expected
