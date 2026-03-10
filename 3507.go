func minimumPairRemoval(nums []int) int {
    var (
        result, minSum, idx int
        sorted bool
    )
    for len(nums) > 1 {
        sorted = true
        minSum = math.MaxInt
        idx = -1
        for i := range len(nums) - 1 {
            if sorted && nums[i] > nums[i+1] {
                sorted = false
            }
            if sum := nums[i] + nums[i+1]; sum < minSum {
                minSum = sum
                idx = i
            }
        }
        if sorted {
            return result
        }

        nums[idx] = minSum
        nums = append(nums[:idx+1], nums[idx+2:]...)
        result++
    }

    return result
}