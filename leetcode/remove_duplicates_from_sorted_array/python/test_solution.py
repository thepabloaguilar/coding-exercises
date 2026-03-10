import pytest

from leetcode.remove_duplicates_from_sorted_array.python.solution import Solution


@pytest.mark.parametrize(
    ('data', 'expected'),
    [
        ([1,1,2], [1,2]),
        ([0,0,1,1,1,2,2,3,3,4], [0,1,2,3,4]),
    ]
)
def test_base_cases(data: list[int], expected: list[int]):
    solution = Solution()

    remaining_nums = solution.removeDuplicates(data)

    assert remaining_nums == len(expected)
    assert data[:remaining_nums] == expected
