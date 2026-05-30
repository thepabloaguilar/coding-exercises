from typing import Any

import pytest

from leetcode.time_based_key_value_store.python.solution import TimeMap


@pytest.mark.parametrize(
    ('data',),
    [
        (
            [
                ('set', ['alice', 'happy', 1], None),
                ('get', ['alice', 1], 'happy'),
                ('get', ['alice', 2], 'happy'),
                ('set', ['alice', 'sad', 3], None),
                ('get', ['alice', 3], 'sad'),
            ],
        ),
        (
            [
                ('set', ['key1', 'value1', 10], None),
                ('get', ['key1', 1], ''),
                ('get', ['key1', 10], 'value1'),
                ('get', ['key1', 11], 'value1'),
            ],
        ),
    ]
)
def test_base_case(data: list[tuple[str, list[Any], str | None]]):
    time_map = TimeMap()

    for command, args, expected in data:
        assert getattr(time_map, command)(*args) == expected
