func containsNearbyDuplicate(nums []int, k int) bool {
    seen := make(map[int]int, len(nums))
    var (
        last int
        ok bool
    )

    for idx, i := range nums {
        last, ok = seen[i]
        if ok && idx - last <= k {
            return true
        }
        seen[i] = idx
    }

    return false
}