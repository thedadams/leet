func myAtoi(s string) int {
    var (
        result int
        max = 2147483647
        min = -2147483648
        factor = 1
        found bool
    )

    for _, c := range s {
        if c == ' ' {
            if found {
                break
            }
            continue
        }

        if c == '-' {
            if found {
                break
            }
            factor = -1
            found = true
            continue
        }
        if c == '+' {
            if found {
                break
            }
            found = true
            continue
        }

        if c < '0' || c > '9' {
            break
        }

        found = true
        result = result * 10 + factor * int(c - '0')
        if result <= min {
            return min
        }
        if result >= max {
            return max
        }
    }

    return result
}