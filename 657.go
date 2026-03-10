func judgeCircle(moves string) bool {
	var x, y int
	for _, c := range moves {
		switch c {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
	}
	return x == 0 && y == 0
}