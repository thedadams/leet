func maxCount(m int, n int, ops [][]int) int {
    for _, op := range ops {
        m = min(op[0], m)
        n = min(op[1], n)
    }

    return m * n
}