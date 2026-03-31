func numJewelsInStones(jewels string, stones string) int {
	j := make(map[rune]struct{}, len(jewels))
	for _, r := range jewels {
		j[r] = struct{}{}
	}

	var (
		result int
		ok     bool
	)
	for _, r := range stones {
		if _, ok = j[r]; ok {
			result++
		}
	}

	return result
}
