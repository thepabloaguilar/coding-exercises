import pytest

from leetcode.lru_cache.python.solution import LRUCache


@pytest.mark.parametrize(
    ('capacity', 'data'),
    [
        (
            2,
            [
                ('put', [1, 1], None),
                ('put', [2, 2], None),
                ('get', [1], 1),
                ('put', [3, 3], None),
                ('get', [2], -1),
                ('put', [4, 4], None),
                ('get', [1], -1),
                ('get', [3], 3),
                ('get', [4], 4),
            ],
        ),
    ],
)
def test_base_cases(capacity: int, data: list[tuple[str, list[int], int | None]]):
    lru_cache = LRUCache(capacity)
    for command, args, expected in data:
        assert getattr(lru_cache, command)(*args) == expected
