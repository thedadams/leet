func isMonotonic(nums []int) bool {
	var decreasing *bool
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			if decreasing == nil {
				decreasing = &[]bool{false}[0]
			} else if *decreasing {
				return false
			}
		} else if nums[i-1] > nums[i] {
			if decreasing == nil {
				decreasing = &[]bool{true}[0]
			} else if !*decreasing {
				return false
			}
		}
	}

	return true
}
