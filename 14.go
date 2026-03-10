func longestCommonPrefix(strs []string) string {
    s := strs[0]
    
    for i := 1; i < len(strs); i++ {
        if s == "" {
            return s
        }

        t := strs[i]
        if len(s) > len(t) {
            tmp := s
            s = t
            t = tmp
        }
        for j := range t {
            if j >= len(s) {
                break
            }
            if s[j] != t[j] {
                s = s[:j]
            }
        }
    }

    return s
}