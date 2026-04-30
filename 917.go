import "unicode"

func reverseOnlyLetters(s string) string {
	start, end := 0, len(s)-1
	result := []byte(s)
	for {
		for !unicode.IsLetter(rune(result[start])) && start < end {
			start++
		}
		for !unicode.IsLetter(rune(result[end])) && start < end {
			end--
		}
		if start >= end {
			return string(result)
		}

		result[start], result[end] = result[end], result[start]
		start++
		end--
	}
}
