func reverseVowels(s string) string {
    start, end := 0, len(s) - 1
    str := strings.Split(s, "")
    vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

    for start < end {
        if slices.Contains(vowels, str[start]) {
            for start < end && !slices.Contains(vowels, str[end]) {
                end--
            }
            if start < end {
                str[start], str[end] = str[end], str[start]
                end--
            }
        }

        start++
    }

    return strings.Join(str, "")
}