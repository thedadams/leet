func reverseWords(s string) string {
    b := make([]byte, 0, len(s))
    for _, w := range strings.Split(s, " ") {
        for i := len(w) - 1; i >= 0; i-- {
            b = append(b, byte(w[i]))
        }

        b = append(b, ' ')
    }

    return string(b[:len(b)-1])
}