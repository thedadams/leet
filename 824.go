import "strings"

func toGoatLatin(sentence string) string {
	var (
		b   strings.Builder
		idx int
	)
	for w := range strings.SplitSeq(sentence, " ") {
		idx++
		switch w[0] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			b.WriteString(w)
		default:
			b.WriteString(w[1:])
			b.WriteByte(w[0])
		}

		b.WriteString("ma")
		b.WriteString(strings.Repeat("a", idx))
		b.WriteString(" ")
	}

	return b.String()[:b.Len()-1]
}
