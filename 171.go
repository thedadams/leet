func titleToNumber(columnTitle string) int {
    var result int
    for _, c := range columnTitle {
        result = 26 * result + int(c - 'A' + 1)
    }

    return result
}