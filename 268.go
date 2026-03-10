func missingNumber(nums []int) int {
    n := len(nums)
    n = ((n + 1) * n) / 2

    for _, i := range nums {
        n -= i
    }

    return n
}