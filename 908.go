func smallestRangeI(nums []int, k int) int {
	minN, maxN := 10000, 0
	for _, n := range nums {
		minN = min(minN, n)
		maxN = max(maxN, n)
	}

	return max(maxN-minN-2*k, 0)
}
