func findErrorNums(nums []int) []int {
    set := make(map[int]struct{}, len(nums))
    for n := range len(nums) {
        set[n+1] = struct{}{}
    }

    result := []int{0, 0}
    for _, num := range nums {
        if _, ok := set[num]; !ok {
            result[0] = num
        } else {
            delete(set, num)
        }
    }

    for n := range set {
        result[1] = n
    }

    return result
}