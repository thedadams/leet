func countSegments(s string) int {
    var (
        result int = 1
        prev rune = ' '
    )

    for _, r := range s {
        if r == ' ' && prev != ' ' {
            result++
        }
        prev = r
    }

    if prev == ' ' {
        result--
    }

    return result
}