func commonChars(words []string) []string {
	letters := make(map[rune]int, 26)
	for _, w := range words {
		lettersThisWord := make(map[rune]rune, 26)
		for _, r := range w {
			letters[r+'a'*lettersThisWord[r]]++
			lettersThisWord[r]++
		}
	}

	result := make([]string, 0, len(letters))
	wordCount := len(words)
	for r, v := range letters {
		if v == wordCount {
			result = append(result, string((r%'a')+'a'))
		}
	}

	return result
}
