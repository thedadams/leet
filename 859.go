func buddyStrings(s, goal string) bool {
	if len(s) != len(goal) || len(s) == 1 {
		return false
	}

	var hasDupeLetter bool
	letters := make(map[byte]struct{}, 26)
	firstIdx, secondIdx := -1, -1
	for i := range s {
		if _, ok := letters[s[i]]; !ok {
			letters[s[i]] = struct{}{}
		} else {
			hasDupeLetter = true
		}

		if s[i] != goal[i] {
			if firstIdx == -1 {
				firstIdx = i
			} else if secondIdx == -1 {
				if s[firstIdx] != goal[i] || s[i] != goal[firstIdx] {
					return false
				}

				secondIdx = i
			} else {
				return false
			}
		}
	}

	return firstIdx != -1 && secondIdx != -1 || firstIdx == -1 && hasDupeLetter
}
