var (
    r2I = map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
)

func romanToInt(s string) int {
    var (
        num int
        base int = r2I[s[len(s) - 1]]
    )

    for i := len(s) - 1; i >= 0; i-- {
        this := r2I[s[i]]
        if this < base {
            num -= this
        } else {
            num += this
            base = this
        }
    }

    return num
}