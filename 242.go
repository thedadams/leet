func isAnagram(s string, t string) bool {
    letters := make(map[rune]int, len(s))

    for _, r := range s {
        letters[r]++
    }

    for _, r := range t {
        letters[r]--
        if letters[r] == 0 {
            delete(letters, r)
        }
    }

    return len(letters) == 0
}