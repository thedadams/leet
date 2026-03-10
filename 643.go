func findMaxAverage(nums []int, k int) float64 {
    var maxAverage, average int
    for i := range k {
        average += nums[i]
    }
    maxAverage = average

    for i := range len(nums) - k {
        average += nums[i+k] - nums[i]
        maxAverage = max(average, maxAverage)
    }

    return float64(maxAverage) / float64(k)
}