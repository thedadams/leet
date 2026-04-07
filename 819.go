import "strings"

func mostCommonWord(paragraph string, banned []string) string {
	b := make(map[string]struct{}, len(banned))
	for _, s := range banned {
		b[s] = struct{}{}
	}

	counts := make(map[string]int, 150)
	var result string
	for _, s := range strings.FieldsFuncSeq(paragraph, func(r rune) bool {
		switch r {
		case ' ', ',', '\'', '!', '?', ';', '.':
			return true
		}
		return false
	}) {
		s = strings.ToLower(s)
		if _, ok := b[s]; !ok {
			counts[s]++
			if counts[s] > counts[result] {
				result = s
			}
		}
	}

	return result
}
