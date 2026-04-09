func shortestToChar(s string, c byte) []int {
	result := make([]int, len(s))
	lastIdx := -1
	for i, r := range []byte(s) {
		result[i] = -1
		if r == c {
			lastIdx = i
			result[i] = 0

			for j := i - 1; j >= 0; j-- {
				if result[j] == -1 || result[j] > i-j {
					result[j] = i - j
				} else {
					break
				}
			}
		} else if lastIdx != -1 {
			result[i] = i - lastIdx
		}
	}

	for j := lastIdx + 1; j < len(result); j++ {
		result[j] = j - lastIdx
	}

	return result
}
