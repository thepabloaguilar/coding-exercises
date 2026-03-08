package golang

func NumEnclaves(grid [][]int) int {
	notConnected := len(grid) * len(grid[0])
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 0 {
				notConnected -= 1
				continue
			}

			isBoundary := i == 0 || i == len(grid)-1 || j == 0 || j == len(grid[0])-1
			if isBoundary {
				notConnected -= dfs(grid, i, j)
			}
		}
	}

	return notConnected
}

func dfs(grid [][]int, i int, j int) int {
	isOutOfBoundaries := i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0])
	if isOutOfBoundaries || grid[i][j] != 1 {
		return 0
	}

	grid[i][j] = 2

	connected := 1
	connected += dfs(grid, i+1, j)
	connected += dfs(grid, i-1, j)
	connected += dfs(grid, i, j+1)
	connected += dfs(grid, i, j-1)

	return connected
}
