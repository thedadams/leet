func addDigits(num int) int {
    var next int
    for num > 9 {
        next = 0
        for num > 0 {
            next += num % 10
            num /= 10
        }

        num = next
    }
    return num
}