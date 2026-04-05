package golang

func CanFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)

	for _, course := range prerequisites {
		graph[course[0]] = append(graph[course[0]], course[1])
	}

	visited := make(map[int]struct{})
	for course := range graph {
		if !dfs(graph, course, visited) {
			return false
		}
	}

	return true
}

func dfs(graph map[int][]int, course int, visited map[int]struct{}) bool {
	if _, ok := visited[course]; ok {
		return false
	}

	if len(graph[course]) == 0 {
		return true
	}

	visited[course] = struct{}{}

	for _, edge := range graph[course] {
		if !dfs(graph, edge, visited) {
			return false
		}
	}

	graph[course] = make([]int, 0)
	delete(visited, course)

	return true
}
