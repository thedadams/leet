func removeOuterParentheses(s string) string {
	openParen := -1
	result := make([]rune, 0, len(s))
	for _, r := range s {
		switch r {
		case '(':
			if openParen > -1 {
				result = append(result, r)
			}
			openParen++
		case ')':
			openParen--
			if openParen > -1 {
				result = append(result, r)
			}
		}
	}

	return string(result)
}
