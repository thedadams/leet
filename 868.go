func binaryGap(n int) int {
	var (
		largestGap  int
		place       int
		previousOne = -1
	)
	for n > 0 {
		if n%2 == 1 {
			if previousOne != -1 {
				largestGap = max(largestGap, place-previousOne)
			}
			previousOne = place
		}
		place++
		n >>= 1
	}

	return largestGap
}
