func isHappy(n int) bool {
    seen := make(map[int]bool)
    for n != 1 && !seen[n] {
        seen[n] = true
        n = sumSquareDigits(n)
    }

    return n == 1
}

func sumSquareDigits(n int) int {
    var sum, digit int
    for n > 0 {
        digit = n % 10
        sum += digit * digit
        n /= 10
    }

    return sum
}