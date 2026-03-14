func findLengthOfLCIS(nums []int) int {
    length, maxLength, prev := 1, 1, nums[0]

    for _, num := range nums {
        if num > prev {
            length++
            maxLength = max(length, maxLength)
        } else {
            length = 1
        }
        prev = num
    }

    return maxLength
}