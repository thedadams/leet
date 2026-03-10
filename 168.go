func convertToTitle(columnNumber int) string {
    var col string
    for columnNumber > 0 {
        columnNumber--
        col = string('A' + columnNumber % 26) + col
        columnNumber /= 26
    }

    return col
}