var letters = []string{
	".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
}

func uniqueMorseRepresentations(words []string) int {
	results := make(map[string]struct{})
	var t string
	for _, w := range words {
		t = ""
		for _, r := range w {
			t += letters[r-'a']
		}

		results[t] = struct{}{}
	}

	return len(results)
}
