import "strings"

func uncommonFromSentences(s1, s2 string) []string {
	var (
		result []string
		m      = make(map[string]int, 45)
	)
	for _, s := range strings.Fields(s1) {
		m[s]++
	}
	for _, s := range strings.Fields(s2) {
		m[s]++
	}

	for s, c := range m {
		if c == 1 {
			result = append(result, s)
		}
	}

	return result
}
