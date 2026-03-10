func hammingDistance(x, y int) int {
    x ^= y

    var result int
    for x > 0 {
        result += x & 1
        x >>= 1
    }

    return result
}