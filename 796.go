func rotateString(s string, goal string) bool {
	for i, r := range s {
		if r == rune(goal[0]) && string(append([]byte(s[i:]), s[:i]...)) == goal {
			return true
		}
	}

	return false
}
