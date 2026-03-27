import "slices"

func areSimilar(mat [][]int, k int) bool {
	k = k % len(mat[0])

	if k != 0 {
		for _, r := range mat {
			if !slices.Equal(r, append(r[k:], r[:k]...)) {
				return false
			}
		}
	}
	return true
}
