class Solution:
    def canFinish(self, numCourses: int, prerequisites: list[list[int]]) -> bool:
        graph = {course: [] for course in range(numCourses)}

        for course, prerequisite in prerequisites:
            graph[course].append(prerequisite)

        visited = set()
        for course in graph:
            if not self._dfs(graph, course, visited):
                return False

        return True

    def _dfs(self, graph: dict[int, list[int]], course: int, visited: set[int]) -> bool:
        if course in visited:
            return False

        if not graph[course]:
            return True

        visited.add(course)

        for edge in graph[course]:
            if not self._dfs(graph, edge, visited):
                return False

        graph[course] = []
        visited.remove(course)

        return True
