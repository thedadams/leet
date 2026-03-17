func countBinarySubstrings(s string) int {
    var firstCount, secondCount, result int
    lookingFor := rune(s[0])
    for _, r := range s {
        if lookingFor == r {
            secondCount++
        } else {
            lookingFor = r
            result += min(firstCount, secondCount)
            firstCount = secondCount
            secondCount = 1
        }
    }

    return result + min(firstCount, secondCount)
}
