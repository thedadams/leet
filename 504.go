func convertToBase7(num int) string {
    if num == 0 {
        return "0"
    }

    neg := num < 0
    if neg {
        num = -num
    }

    pow := 1
    for pow <= num {
        pow *= 7
    }
    pow /= 7

    var result string
    for pow > 0 {
        result += strconv.Itoa(num / pow)
        num %= pow
        pow /= 7
    }

    if neg {
        return "-"+ result
    }

    return result
}