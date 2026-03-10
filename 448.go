func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); {
		if nums[i] != i+1 && nums[i] != nums[nums[i] - 1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

    var result []int
    for i := range len(nums) {
        if nums[i] != i + 1 {
            result = append(result, i + 1)
        }
    }
	return result
}