func wordPattern(pattern string, s string) bool {
    toPattern := make(map[string]byte, len(pattern))
    toString := make(map[byte]string, len(pattern))

    fields := strings.Split(s, " ")
    if len(fields) != len(pattern) {
        return false
    }

    for i, t := range fields {
        b, patternExists := toPattern[t]
        u, stringExists := toString[pattern[i]]
        if patternExists != stringExists || (stringExists && t != u) || (patternExists && b != pattern[i]) {
            return false
        }

        toPattern[t] = pattern[i]
        toString[pattern[i]] = t
    }

    return true
}