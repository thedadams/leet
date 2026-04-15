func transpose(matrix [][]int) [][]int {
	result := make([][]int, len(matrix[0]))
	for j := range len(result) {
		result[j] = make([]int, len(matrix))
	}
	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}
