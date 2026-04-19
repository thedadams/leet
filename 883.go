func projectionArea(grid [][]int) int {
	var result int
	maxZs := make([]int, len(grid[0]))
	for _, y := range grid {
		var maxZ int
		for i, z := range y {
			if z > 0 {
				// x-y
				result++
			}
			// y-z
			maxZ = max(maxZ, z)
			maxZs[i] = max(maxZs[i], z)
		}
		result += maxZ
	}

	for _, maxZ := range maxZs {
		// x-z
		result += maxZ
	}

	return result
}
