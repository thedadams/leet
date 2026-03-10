func maximumDifference(nums []int) int {
    var (
        minSoFar = nums[0]
        maxDiff int = -1
    )

    for _, i := range nums {
        if i > minSoFar {
            maxDiff = max(maxDiff, i - minSoFar)
        } else {
            minSoFar = i
        }
    }

    return maxDiff
}