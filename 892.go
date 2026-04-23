import "math"

func surfaceArea(grid [][]int) int {
	result := 2 * len(grid) * len(grid)
	for i, row := range grid {
		for j, item := range row {
			if item == 0 {
				// No surface area, but we counted it on initialization above
				result -= 2
			}

			if i == 0 || i == len(grid)-1 {
				// Outer faces
				result += item
				if i == 0 && i == len(grid)-1 {
					// Double count because we only go through this loop once
					result += item
				}
			}

			if j == 0 || j == len(row)-1 {
				// More outer faces
				result += item
				if i == 0 && i == len(grid)-1 {
					// Double count because we only go through this loop once
					result += item
				}
			}

			if j > 0 {
				// Previous face left
				result += int(math.Abs(float64(row[j-1] - item)))
			}

			if i > 0 {
				// Previous face down
				result += int(math.Abs(float64(grid[i-1][j] - item)))
			}
		}
	}

	return result
}
