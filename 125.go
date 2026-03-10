func isPalindrome(s string) bool {
	s = strings.ToLower(strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return -1
		}
		return r
	}, s))
	l := len(s)
	for i := range l / 2 {
		if s[i] != s[l-i-1] {
			fmt.Printf("%s != %s\n", string(s[i]), string(s[l-i-1]))
			return false
		}
	}

	return true
}