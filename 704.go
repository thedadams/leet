func search(nums []int, target int) int {
    start, end := 0, len(nums) - 1

    mid, num := (end + start) / 2, 0
    for start < end {
        num = nums[mid]
        if num < target {
            start = mid + 1
        } else if num > target {
            end = mid - 1
        } else {
            break
        }
        mid = (end + start) / 2
    }

    if nums[mid] == target {
        return mid
    }
    return -1
}
