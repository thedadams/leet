import "slices"

func diStringMatch(s string) []int {
	result := make([]int, len(s)+1)
	for i := range len(result) {
		result[i] = i
	}

	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			continue
		}

		j := i + 1
		for j < len(s) && s[j] == 'D' {
			j++
		}

		slices.Reverse(result[i : j+1])
		i = j
	}

	return result
}

