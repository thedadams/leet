func bitwiseComplement(n int) int {
    if n == 0 {
        return 1
    }
    var result, i int
    for n > 0 {
        if n % 2 == 0 {
            result += 1<<i
        }
        n >>= 1
        i++
    }

    return result
}