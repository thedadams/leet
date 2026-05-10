func minDeletionSize(strs []string) int {
	var result int

	for i := range len(strs[0]) {
		for j := range len(strs) - 1 {
            if strs[j][i] > strs[j+1][i] {
                result++
                break
            }
		}
	}

    return result
}
