func selfDividingNumbers(left int, right int) []int {
    var result []int
    for i := left; i <= right; i++ {
        if isSelfDividing(i) {
            result = append(result, i)
        }
    }

    return result
}

func isSelfDividing(n int) bool {
    var d int
    m := n
    for m > 0 {
        d = m % 10
        if d == 0 || n % d != 0 {
            return false
        }
        m /= 10
    }

    return true
}
