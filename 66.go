func plusOne(digits []int) []int {
    i := 1
    j := len(digits) - 1
    for i == 1 {
        if j == -1 {
            digits = append(digits, 0)
            for k := 1; k < len(digits); k++ {
                digits[k-1] = digits[k]
            }
            j = 0
        }

        digits[j] += i
        if digits[j] == 10 {
            digits[j] = 0
            j--
        } else {
            i = 0
        }
    }

    return digits
}