func findJudge(n int, trust [][]int) int {
	trusts := make([]int, n)
	for _, t := range trust {
		trusts[t[0]-1]--
		trusts[t[1]-1]++
	}

	for i, s := range trusts {
		if s == n-1 {
			return i + 1
		}
	}

	return -1
}
