func dominantIndex(nums []int) int {
	var maxIdx, secondMax int
	for i, n := range nums {
		if n > nums[maxIdx] {
			secondMax = nums[maxIdx]
			maxIdx = i
		} else if i > 0 && n > secondMax {
			secondMax = n
		}
	}

	if nums[maxIdx] >= 2*secondMax {
		return maxIdx
	}
	return -1
}
