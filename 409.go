func longestPalindrome(s string) int {
    seen := make(map[rune]int)
    
    var ans int
    for _, r := range s {
        seen[r]++
        if seen[r] % 2 == 0 {
            ans += 2
        }
    }

    for _, count := range seen {
        if count % 2 == 1 {
            ans++
            break
        }
    }

    return ans
}