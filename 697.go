func findShortestSubArray(nums []int) int {
	var (
		firstLastCount = make(map[int][2]int, len(nums))
		result         = math.MaxInt
		maxCount       int
		s              [2]int
		ok             bool
	)
	for i, n := range nums {
		s, ok = firstLastCount[n]
		if !ok {
			s[0] = i
		}
		s[1]++
		firstLastCount[n] = s

		if maxCount < s[1] {
			result = i - s[0] + 1
			maxCount = s[1]
		} else if maxCount == s[1] {
			result = min(result, i-s[0]+1)
		}
	}

	return result
}
