import pytest

from leetcode.best_time_to_buy_and_sell_stock_ii.python.solution import Solution


@pytest.mark.parametrize(
    ('prices', 'expected'),
    (
        ([7,1,5,3,6,4], 7),
        ([1,2,3,4,5], 4),
        ([7,6,4,3,1], 0),
    ),
)
def test_base_cases(prices: list[int], expected: int):
    solution = Solution()
    assert solution.maxProfit(prices) == expected
