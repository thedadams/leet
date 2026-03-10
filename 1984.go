func minimumDifference(nums []int, k int) int {
    if k == 1 {
        return 0
    }

    slices.Sort(nums)
    
    result := math.MaxInt
    for i := k-1; i < len(nums); i++ {
        result = min(result, nums[i] - nums[i-k+1])
    }

    return result
}