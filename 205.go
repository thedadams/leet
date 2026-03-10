func isIsomorphic(s, t string) bool {
    translation := make(map[rune]byte, len(s))
    for i, p := range s {
        q := t[i]
        if r, ok := translation[p]; !ok {
            translation[p] = q
        } else if r != q {
            return false
        }
    }

    translation = make(map[rune]byte, len(t))
    for i, p := range t {
        q := s[i]
        if r, ok := translation[p]; !ok {
            translation[p] = q
        } else if r != q {
            return false
        }
    }

    return true
}