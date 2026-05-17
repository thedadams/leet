func repeatedNTimes(nums []int) int {
    var idx int
    for idx + 2 < len(nums) {
        if nums[idx] == nums[idx+1] || nums[idx] == nums[idx+2] {
            return nums[idx]
        }
        if nums[idx+1] == nums[idx+2] {
            return nums[idx+1]
        }
        idx += 3
    }

    return nums[len(nums)-1]
}
