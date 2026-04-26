func sortArrayByParity(nums []int) []int {
	end := len(nums) - 1
	for i := 0; i < end; i++ {
		if nums[i]%2 == 1 {
			nums[i], nums[end] = nums[end], nums[i]
			end--
			i--
		}
	}

	return nums
}
