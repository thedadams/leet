import "strings"

func shortestCompletingWord(licensePlate string, words []string) string {
	l := make(map[rune]int)
	for _, r := range strings.ToLower(licensePlate) {
		if r >= 'a' && r <= 'z' {
			l[r]++
		}
	}

	var (
		result string
		m      = make(map[rune]int)
		fits   bool
	)
	for _, w := range words {
		clear(m)
		fits = true
		for _, r := range w {
			m[r]++
		}

		for k, v := range l {
			if v > m[k] {
				fits = false
				break
			}
		}

		if fits && (len(w) < len(result) || len(result) == 0) {
			result = w
		}
	}
	return result
}
