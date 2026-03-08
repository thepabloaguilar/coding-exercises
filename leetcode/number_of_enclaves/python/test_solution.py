import pytest

from leetcode.number_of_enclaves.python.solution import Solution

@pytest.mark.parametrize(
    ('data', 'expected'),
    [
        (
                [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]],
                3,
        ),
        (
                [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]],
                0,
        ),
    ],
)
def test_base_cases(data: list[list[int]], expected: int) -> None:
    solution = Solution()
    assert solution.numEnclaves(data) == expected
