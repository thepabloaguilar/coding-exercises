import pytest

from leetcode.h_index.python.solution import Solution


@pytest.mark.parametrize(
    ('citations', 'expected'),
    [
        ([3,0,6,1,5], 3),
        ([1,3,1], 1),
        ([100], 1),
    ]
)
def test_base_cases(citations: list[int], expected: int):
    solution = Solution()
    assert solution.hIndex(citations) == expected
