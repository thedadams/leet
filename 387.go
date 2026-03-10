func firstUniqChar(s string) int {
    letters := make(map[rune]int)
    
    for _, r := range s {
        letters[r]++
    }

    for i, r := range s {
        if letters[r] == 1 {
            return i
        }
    }

    return -1
}