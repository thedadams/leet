func backspaceCompare(s, t string) bool {
	return replace(s) == replace(t)
}

func replace(s string) string {
	b := make([]rune, 0, len(s))
	for _, r := range s {
		if r == '#' {
			if len(b) > 0 {
				b = b[:len(b)-1]
			}
		} else {
			b = append(b, r)
		}
	}

	return string(b)
}
