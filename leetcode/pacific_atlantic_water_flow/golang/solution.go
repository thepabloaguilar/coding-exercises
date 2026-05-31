package golang

func PacificAtlantic(heights [][]int) [][]int {
	rowLen, columnLen := len(heights), len(heights[0])
	if rowLen == 0 || columnLen == 0 {
		return nil
	}

	pacificReachable := make([][]bool, rowLen)
	for idx := range pacificReachable {
		pacificReachable[idx] = make([]bool, columnLen)
	}

	atlanticReachable := make([][]bool, rowLen)
	for idx := range atlanticReachable {
		atlanticReachable[idx] = make([]bool, columnLen)
	}

	// Check Pacific Ocean reachability
	for idx := range rowLen {
		dfs(heights, idx, 0, pacificReachable)
	}

	for idx := range columnLen {
		dfs(heights, 0, idx, pacificReachable)
	}

	// Check Atlantic Ocean reachability
	for idx := range rowLen {
		dfs(heights, idx, columnLen-1, atlanticReachable)
	}

	for idx := range columnLen {
		dfs(heights, rowLen-1, idx, atlanticReachable)
	}

	// Check overall reachability
	result := make([][]int, 0)
	for row := range rowLen {
		for column := range columnLen {
			if pacificReachable[row][column] && atlanticReachable[row][column] {
				result = append(result, []int{row, column})
			}
		}
	}

	return result
}

func dfs(heights [][]int, row int, column int, reachable [][]bool) {
	reachable[row][column] = true

	directions := [4][2]int{
		{1, 0},  // Up
		{-1, 0}, // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}
	for _, direction := range directions {
		newRow, newColumn := row+direction[0], column+direction[1]

		isWithinBoundaries := newRow >= 0 && newRow < len(heights) && newColumn >= 0 && newColumn < len(heights[0])
		if isWithinBoundaries && !reachable[newRow][newColumn] {
			if heights[newRow][newColumn] >= heights[row][column] {
				dfs(heights, newRow, newColumn, reachable)
			}
		}
	}
}
