func canConstruct(ransomNote string, magazine string) bool {
    intersection := make(map[rune]int, len(ransomNote))
    for _, r := range magazine {
        intersection[r]++
    }

    for _, r := range ransomNote {
        if intersection[r] == 0 {
            return false
        }
        intersection[r]--
    }

    return true
}