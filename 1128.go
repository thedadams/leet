func numEquivDominoPairs(dominoes [][]int) int {
    seenDominoes := make(map[[2]int]int, len(dominoes))
    
    var (
        count int
        d [2]int
    )
    for _, dom := range dominoes {
        d = [2]int{dom[0], dom[1]}
        if d[0] > d[1] {
            d[0], d[1] = d[1], d[0]
        }

        count += seenDominoes[d]
        seenDominoes[d]++
    }

    return count
}