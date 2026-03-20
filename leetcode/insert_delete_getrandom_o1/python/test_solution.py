import random

import pytest

from leetcode.insert_delete_getrandom_o1.python.solution import RandomizedSet


@pytest.mark.parametrize(
    ('data',),
    [
        (
            [
                ('insert', [1], True),
                ('remove', [2], False),
                ('insert', [2], True),
                ('getRandom', [], 2),
                ('remove', [1], True),
                ('insert', [2], False),
                ('getRandom', [], 2),
            ],
        ),
        (
            [
                ('insert', [1], True),
                ('remove', [1], True),
            ],
        ),
    ],
)
def test_base_cases(data: list[tuple[str, list[int], bool]]):
    random.seed(0)

    randomized_set = RandomizedSet()
    for command, args, expected in data:
        assert getattr(randomized_set, command)(*args) == expected
