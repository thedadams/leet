func toHex(num int) string {
    num64 := int64(num)
    if num64 < 0 {
        num64 = twosComplement(num64)
    } else if num == 0 {
        return "0"
    }

    var result string
    var powerOf16 int64 = 1
    for powerOf16 <= num64 {
        powerOf16 *= 16
    }

    powerOf16 /= 16

    var digit int64
    for powerOf16 > 0 {
        digit = num64 / powerOf16
        if digit > 9 {
            result += fmt.Sprintf("%c", int64('a') + digit - 10)
        } else {
            result += fmt.Sprintf("%d", digit)
        }

        num64 %= powerOf16
        powerOf16 /= 16
    }

    return result
}

func twosComplement(num int64) int64 {
    return ((1<< 32 - 1) ^ -num) + 1
}