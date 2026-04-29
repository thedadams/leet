func hasGroupsSizeX(deck []int) bool {
	var counts [10000]int
	for _, c := range deck {
		counts[c]++
	}

	minCount := -1
	for _, c := range counts {
		if c > 0 {
			minCount = gcd(minCount, c)
		}
	}

	return minCount != 1
}

func gcd(n, m int) int {
	if n == -1 {
		return m
	}
	if n > m {
		n, m = m, n
	}

	for r := m % n; r > 0; r = m % n {
		m = n
		n = r
	}
	return n
}
