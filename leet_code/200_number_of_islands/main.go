func numIslands(grid [][]byte) int {
	// obtain grid dimensions
	x := len(grid[0])
	y := len(grid)

	// declare count variable
	var count int

	// declare backtrack func variable
	var backtrack func(j, i int)

	backtrack = func(j, i int) {
		// start backtracking in all four directions,
		// but first make sure indexes are not out of boundaries (return in that case)
		if j < 0 || j >= y || i < 0 || i >= x || grid[j][i] != '1' {
			return
		}

		// we change the point in the grid to something else, so
		// it is not considered in future backtracking calls
		grid[j][i] = '#'

		backtrack(j+1, i)
		backtrack(j-1, i)
		backtrack(j, i+1)
		backtrack(j, i-1)
	}

	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			// only start back track when there is a '1'
			if grid[j][i] == '1' {
				// grid[j][i] = '#'
				count++
				backtrack(j, i)
			}
		}
	}
	return count
}