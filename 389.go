func findTheDifference(s, t string) byte {
    var char rune
    for _, r := range s {
        char ^= r
    }

    for _, r := range t {
        char ^= r
    }

    return byte(char)
}