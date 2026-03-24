package golang

func NumIslands(grid [][]byte) int {
	islands := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				traverse(grid, i, j)
				islands++
			}
		}
	}

	return islands
}

func traverse(grid [][]byte, i int, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	traverse(grid, i+1, j)
	traverse(grid, i-1, j)
	traverse(grid, i, j+1)
	traverse(grid, i, j-1)
}
