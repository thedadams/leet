func containsDuplicate(nums []int) bool {
    seen := make(map[int]struct{})
    var ok bool
    for _, i := range nums {
        if _, ok = seen[i]; ok {
            return true
        }
        seen[i] = struct{}{}
    }

    return false
}