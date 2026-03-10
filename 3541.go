func maxFreqSum(s string) int {
    vowels := map[rune]int{
        'a': 0,
        'e': 0,
        'i': 0,
        'o': 0,
        'u': 0,
    }

    consonants := make(map[rune]int, 21)

    var m, n int
    for _, r := range s {
        _, ok := vowels[r]
        if ok {
            vowels[r]++
            m = max(m, vowels[r])
        } else {
            consonants[r]++
            n = max(n, consonants[r])
        }
    }

    return m + n
}