func isToeplitzMatrix(matrix [][]int) bool {
	for i := range matrix {
		num := matrix[i][0]
		for j := 0; i+j < len(matrix) && j < len(matrix[i+j]); j++ {
			if matrix[i+j][j] != num {
				return false
			}
		}
	}

	for j := range matrix[0] {
		num := matrix[0][j]
		for i := 0; i < len(matrix) && i+j < len(matrix[i]); i++ {
			if matrix[i][i+j] != num {
				return false
			}
		}
	}

	return true
}
