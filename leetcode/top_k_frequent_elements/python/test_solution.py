import pytest

from leetcode.top_k_frequent_elements.python.solution import Solution


@pytest.mark.parametrize(
    ('nums', 'k', 'expected'),
    [
        ([1, 1, 1, 2, 2, 3], 2, [1, 2]),
        ([1], 1, [1]),
        ([1, 2, 1, 2, 1, 2, 3, 1, 3, 2], 2, [1, 2]),
    ]
)
def test_base_cases(nums: list[int], k: int, expected: list[int]):
    solution = Solution()
    assert solution.topKFrequent(nums, k) == expected
