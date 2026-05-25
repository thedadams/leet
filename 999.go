func numRookCaptures(board [][]byte) int {
	var rookI, rookJ int
outer:
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'R' {
				rookI, rookJ = i, j
				break outer
			}
		}
	}

	var result int
	i := rookI - 1
	for ; i >= 0 && board[i][rookJ] == '.'; i-- {
	}
	if i >= 0 && board[i][rookJ] == 'p' {
		result++
	}

	i = rookI + 1
	for ; i < len(board) && board[i][rookJ] == '.'; i++ {
	}
	if i < len(board) && board[i][rookJ] == 'p' {
		result++
	}

	j := rookJ - 1
	for ; j >= 0 && board[rookI][j] == '.'; j-- {
	}
	if j >= 0 && board[rookI][j] == 'p' {
		result++
	}

	j = rookJ + 1
	for ; j < len(board[rookI]) && board[rookI][j] == '.'; j++ {
	}
	if j < len(board[rookI]) && board[rookI][j] == 'p' {
		result++
	}

	return result
}
