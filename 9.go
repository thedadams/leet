func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    n := 1
    for 10*n <= x {
        n *= 10
    }
    for i := 10; i <= n; i, n = i * 10, n / 10 {
        if (x / (i / 10)) % 10 != (x / n) % 10 {
            return false
        }
    }
    
    return true
}