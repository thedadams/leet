import "slices"

func isAlienSorted(words []string, order string) bool {
	alphabet := make(map[byte]byte, len(order))
	for i, c := range order {
		alphabet[byte(c)] = byte(i)
	}

	for i, s := range words {
		word := make([]byte, len(s))
		for i, c := range s {
			word[i] = 'a' + alphabet[byte(c)]
		}
		words[i] = string(word)
	}

	return slices.IsSorted(words)
}
