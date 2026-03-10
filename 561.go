func arrayPairSum(nums []int) int {
    counts := make(map[int]int, len(nums))

    var minimum, maximum = -10000, 10000
    for _, i := range nums{
        counts[i]++
        minimum = min(minimum, i)
        maximum = max(maximum, i)
    }

    var result, count int
    for k := minimum; k <= maximum; k++ {
        for range counts[k] {
            if count % 2 == 0 {
                result += k
            }
            count++
        }
    }

    return result
}