import pytest

from leetcode.gas_station.python.solution import Solution


@pytest.mark.parametrize(
    ('gas', 'cost', 'expected'),
    (
        ([1,2,3,4,5], [3,4,5,1,2], 3),
        ([2,3,4], [3,4,3], -1),
        ([3,3,4], [3,4,4], -1),
        ([5,1,2,3,4], [4,4,1,5,1], 4),
        ([5,8,2,8], [6,5,6,6], 3),
    ),
)
def test_base_cases(gas: list[int], cost: list[int], expected: int):
    solution = Solution()
    assert solution.canCompleteCircuit(gas, cost) == expected
