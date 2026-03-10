func checkOnesSegment(s string) bool {
    var i int
    for i < len(s) && s[i] == '1' {
        i++
    }

    for i < len(s) && s[i] == '0' {
        i++
    }

    return i == len(s)
}