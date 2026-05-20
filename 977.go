func sortedSquares(nums []int) []int {
	idx, neg, pos := 0, 0, 0
	result := make([]int, len(nums))

	for pos < len(nums) && nums[pos] < 0 {
		pos++
	}
	neg = pos - 1

	for idx < len(nums) {
		if neg >= 0 && (pos >= len(nums) || -nums[neg] < nums[pos]) {
			result[idx] = nums[neg] * nums[neg]
			neg--
		} else {
			result[idx] = nums[pos] * nums[pos]
			pos++
		}
		idx++
	}

	return result
}
