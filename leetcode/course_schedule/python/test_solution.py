import pytest

from leetcode.course_schedule.python.solution import Solution


@pytest.mark.parametrize(
    ('numCourses', 'prerequisites', 'expected'),
    (
        (2, [[1, 0]], True),
        (2, [[1, 0], [0, 1]], False),
    ),
)
def test_base_cases(numCourses: int, prerequisites: list[list[int]], expected: int):
    solution = Solution()
    assert solution.canFinish(numCourses, prerequisites) == expected
