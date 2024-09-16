package src

import "sort"

func generate2DArray(rows, cols, defaultValue int) [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
		for j := range grid[i] {
			grid[i][j] = defaultValue
		}
	}
	return grid
}

func clone(grid [][]int) [][]int {
	rows := len(grid)
	newGrid := make([][]int, rows)
	for i := range grid {
		newGrid[i] = append([]int(nil), grid[i]...)
	}
	return newGrid
}

func floodFill(grid [][]int, row, col, value int) {
	numRows := len(grid)
	numCols := len(grid[0])
	originalColor := grid[row][col]

	visited := make([][]bool, numRows)
	for i := range visited {
		visited[i] = make([]bool, numCols)
	}

	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= numRows || c < 0 || c >= numCols || grid[r][c] != originalColor || visited[r][c] {
			return
		}
		grid[r][c] = value
		visited[r][c] = true

		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	dfs(row, col)
}

func calculateConnectedRegion(grid [][]int, row, col int) int {
	numRows := len(grid)
	numCols := len(grid[0])
	color := grid[row][col]

	stack := []struct{ r, c int }{{row, col}}
	visited := make([][]bool, numRows)
	for i := range visited {
		visited[i] = make([]bool, numCols)
	}

	visited[row][col] = true
	regionSize := 0

	dirs := []struct{ dr, dc int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		regionSize++

		for _, dir := range dirs {
			newRow := current.r + dir.dr
			newCol := current.c + dir.dc
			if newRow >= 0 && newRow < numRows && newCol >= 0 && newCol < numCols && !visited[newRow][newCol] && grid[newRow][newCol] == color {
				visited[newRow][newCol] = true
				stack = append(stack, struct{ r, c int }{newRow, newCol})
			}
		}
	}

	return regionSize
}

func evaluate(grid [][]int, colors []int, row, col int) (int, int) {
	numRows := len(grid)
	numCols := len(grid[0])
	color := grid[row][col]

	visited := make([][]bool, numRows)
	for i := range visited {
		visited[i] = make([]bool, numCols)
	}

	regionSize := 0

	count := make([]int, len(colors))

	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= numRows || c < 0 || c >= numCols {
			return
		}
		if visited[r][c] {
			return
		}

		visited[r][c] = true

		if grid[r][c] != color {
			count[grid[r][c]-1]++
			return
		}

		regionSize++

		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	dfs(row, col)

	// Sort by growth descending using Go's efficient sort package
	sort.Slice(count, func(i, j int) bool {
		return count[i] > count[j]
	})

	return regionSize, count[0] + count[1]
}
