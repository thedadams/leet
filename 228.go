func summaryRanges(nums []int) []string {
    if len(nums) == 0 {
        return nil
    }

    var ranges []string
    startRange := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] - nums[i-1] > 1 {
            if startRange == nums[i-1] {
                ranges = append(ranges, strconv.Itoa(startRange))
            } else {
                ranges = append(ranges, fmt.Sprintf("%d->%d", startRange, nums[i-1]))
            }

            startRange = nums[i]
        }
    }

    if startRange == nums[len(nums) - 1] {
        ranges = append(ranges, strconv.Itoa(startRange))
    } else {
        ranges = append(ranges, fmt.Sprintf("%d->%d", startRange, nums[len(nums) - 1]))
    }

    return ranges
}