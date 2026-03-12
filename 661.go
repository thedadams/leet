func imageSmoother(img [][]int) [][]int {
    result := make([][]int, len(img))
    for i := range img {
        result[i] = make([]int, len(img[i]))
        for j := range img[i] {
            result[i][j] = smoothIndex(img, i, j)
        }
    }

    return result
}

func smoothIndex(img [][]int, i, j int) int {
    var result, count int
    for _, iDiff := range []int{-1, 0, 1} {
        for _, jDiff := range []int{-1, 0, 1} {
            newI, newJ := i + iDiff, j +jDiff
            if newI < len(img) && newI >= 0 && newJ < len(img[i]) && newJ >= 0 {
                result += img[newI][newJ]
                count++
            }
        }
    }

    return result / count
}
