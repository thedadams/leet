func reverseStr(s string, k int) string {
    var (
        t []byte
        reverse bool = true
    )
    for i := 0; i < len(s); i += k {
        m := min(len(s), i + k)
        if reverse {
            for j := m-1; j >= i; j-- {
                t = append(t, s[j])
            }
        } else {
            t = append(t, s[i:m]...)
        }
        reverse = !reverse
    }

    return string(t)
}