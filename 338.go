func countBits(n int) []int {
    ones := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        ones[i] = ones[i >> 1] + (i & 1)
    }

    return ones
}