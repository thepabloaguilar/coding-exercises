package golang

func FindCircleNum(isConnected [][]int) int {
	numberOfProvinces := 0
	visited := make([]bool, len(isConnected))

	for idx := range isConnected {
		if !visited[idx] {
			numberOfProvinces++
			dfs(isConnected, visited, idx)
		}
	}

	return numberOfProvinces
}

func dfs(isConnected [][]int, visited []bool, nodeIdx int) {
	visited[nodeIdx] = true
	for nextNodeIdx, value := range isConnected[nodeIdx] {
		if value == 1 && !visited[nextNodeIdx] {
			dfs(isConnected, visited, nextNodeIdx)
		}
	}
}
