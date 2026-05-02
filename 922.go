func sortArrayByParityII(nums []int) []int {
	j := 1
	for i := 0; i < len(nums); i += 2 {
		if nums[i]%2 != i%2 {
			for ; j < len(nums); j += 2 {
				if nums[i]%2 != nums[j]%2 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}

	return nums
}
