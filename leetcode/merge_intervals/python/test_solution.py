import pytest

from leetcode.merge_intervals.python.solution import Solution


@pytest.mark.parametrize(
    ('intervals', 'expected'),
    [
        (
            [[1, 3], [2, 6], [8, 10], [15, 18]],
            [[1, 6], [8, 10], [15, 18]],
        ),
        (
            [[1, 4], [4, 5]],
            [[1, 5]],
        ),
        (
            [[4, 7], [1, 4]],
            [[1, 7]],
        ),
    ]
)
def test_base_cases(intervals: list[list[int]], expected: list[list[int]]):
    solution = Solution()
    assert solution.merge(intervals) == expected
