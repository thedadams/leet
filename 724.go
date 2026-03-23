func pivotIndex(nums []int) int {
	var rightSum int
	for _, n := range nums {
		rightSum += n
	}

	var leftSum int
	for i, n := range nums {
		rightSum -= n
		if leftSum == rightSum {
			return i
		}
		leftSum += n
	}

	return -1
}
