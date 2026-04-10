func largeGroupPositions(s string) [][]int {
	var (
		result [][]int
		idx    int
		char   rune
	)
	for i, r := range s {
		if r != char {
			if i-idx > 2 {
				result = append(result, []int{idx, i - 1})
			}
			idx = i
			char = r
		}
	}

	if len(s)-idx > 2 {
		result = append(result, []int{idx, len(s) - 1})
	}

	return result
}
