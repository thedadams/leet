func licenseKeyFormatting(s string, k int) string {
    s = strings.ToUpper(strings.ReplaceAll(s, "-", ""))
    result := make([]byte, len(s) + len(s) / k)

    groupSize := k
    pos := len(result) - 1
    for i := len(s) - 1; i >= 0; i-- {
        result[pos] = byte(s[i])
        pos--
        groupSize--
        if groupSize == 0 && i != 0 {
            result[pos] = '-'
            pos--
            groupSize = k
        }
    }

    return string(result[pos+1:])
}