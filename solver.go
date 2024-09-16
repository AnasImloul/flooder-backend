package src

import (
	"sort"
)

func greedyFill(grid [][]int, colors []int) int {
	rows := len(grid)
	cols := len(grid[0])
	currentColor := grid[0][0]
	rounds := 0
	regionSize := calculateConnectedRegion(grid, 0, 0)

	for regionSize < rows*cols {
		maxGrowth := 0
		maxEval := 0
		bestColor := currentColor
		bestGrid := clone(grid)

		for _, color := range colors {
			if color != currentColor {
				newGrid := clone(grid)
				floodFill(newGrid, 0, 0, color)

				newRegionSize, eval := evaluate(newGrid, colors, 0, 0)

				if newRegionSize > regionSize && newRegionSize+eval > maxEval {
					maxGrowth = newRegionSize
					maxEval = eval + newRegionSize
					bestColor = color
					bestGrid = newGrid
				}
			}
		}

		grid = bestGrid
		currentColor = bestColor
		regionSize = maxGrowth
		rounds++
	}

	return rounds
}

func findMinRoundsWithGrowth(grid [][]int, colors []int, iterations int) int {
	rows := len(grid)
	cols := len(grid[0])
	minRounds := greedyFill(grid, colors)

	totalSimulations := iterations

	var helper func([][]int, int, int) int
	helper = func(grid [][]int, rounds, regionSize int) int {
		if totalSimulations <= 0 {
			return minRounds
		}

		if regionSize == rows*cols || rounds >= minRounds {
			minRounds = min(minRounds, rounds)
			totalSimulations--
			return minRounds
		}

		currentColor := grid[0][0]

		colorGrowths := make([]struct {
			grid     [][]int
			color    int
			growth   int
			evaluate int
		}, 0, len(colors))

		for _, color := range colors {
			if color != currentColor {

				newGrid := clone(grid)
				floodFill(newGrid, 0, 0, color)
				newRegionSize, eval := evaluate(newGrid, colors, 0, 0)

				if newRegionSize == rows*cols {
					minRounds = min(minRounds, rounds+1)
					totalSimulations--
					return minRounds
				}
				if newRegionSize > regionSize {
					colorGrowths = append(colorGrowths, struct {
						grid     [][]int
						color    int
						growth   int
						evaluate int
					}{grid: newGrid, color: color, growth: newRegionSize, evaluate: eval})
				}
			}
		}

		// Sort by growth descending using Go's efficient sort package
		sort.Slice(colorGrowths, func(i, j int) bool {
			return colorGrowths[i].growth+colorGrowths[i].evaluate > colorGrowths[j].growth+colorGrowths[j].evaluate
		})

		// Prune the bottom moves (explore only top 2 growths at most)
		i := 0
		for i < min(len(colorGrowths), 2) {
			helper(colorGrowths[i].grid, rounds+1, colorGrowths[i].growth)
			i++
		}

		return minRounds
	}

	helper(grid, 0, calculateConnectedRegion(clone(grid), 0, 0))
	return minRounds
}
