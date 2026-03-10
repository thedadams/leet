func findLHS(nums []int) int {
    m := make(map[int]int, len(nums))
    for _, n := range nums {
        m[n]++
    }

    var result int
    for n := range m {
        if l := m[n+1]; l > 0 {
            result = max(result, l + m[n])
        }
    }

    return result
}