func hasAlternatingBits(n int) bool {
    bit := n % 2
    n >>= 1
    for n > 0 {
        if bit == n % 2 {
            return false
        }
        bit ^= 1
        n >>= 1
    }

    return true
}
