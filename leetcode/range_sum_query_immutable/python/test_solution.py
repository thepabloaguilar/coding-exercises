import pytest

from leetcode.range_sum_query_immutable.python.solution import NumArray


@pytest.mark.parametrize(
    ('nums', 'query'),
    (
        (
            [-2,0,3,-5,2,-1],
            [
                ([0,2], 1),
                ([2,5], -1),
                ([0,5], -3),
            ],
        ),
    ),
)
def test_base_cases(nums: list[int], query: list[tuple[list[int], int]]):
    nums_array = NumArray(nums)
    for q in query:
        assert nums_array.sumRange(*q[0]) == q[1]
