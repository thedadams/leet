func findLUSlength(a, b string) int {
    if a != b {
        return max(len(a), len(b))
    }

    return -1
}