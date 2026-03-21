func toLowerCase(s string) string {
    result := make([]byte, 0, len(s))
    for _, r := range s {
        if r >= 'A' && r <= 'Z' {
            r = 'a' + r - 'A'
        }
        result = append(result , byte(r))
    }

    return string(result)
}
