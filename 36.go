func isValidSudoku(board [][]byte) bool {
	var (
		colSeen    [9][9]bool
		squareSeen [3][9]bool
		seen       [9]bool
		el         byte
	)
	for i := range 9 {
		seen = [9]bool{}
		if i%3 == 0 {
			squareSeen = [3][9]bool{}
		}

		for j := range 9 {
			el = board[i][j]
			if el == '.' {
				continue
			}

            el -= '1'

			if seen[el] {
				return false
			}
			seen[el] = true

			if colSeen[j][el] {
				return false
			}
			colSeen[j][el] = true

			if squareSeen[j/3][el] {
				return false
			}
			squareSeen[j/3][el] = true
		}
	}

	return true
}