func findContentChildren(g, s []int) int {
    slices.Sort(g)
    slices.Sort(s)

    var ans, j int
    for i := range len(g) {
        for j < len(s) && s[j] < g[i] {
            j++
        }

        if j < len(s) {
            ans++
        }

        j++
    }
    return ans
}