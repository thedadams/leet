func matrixReshape(mat [][]int, r, c int) [][]int {
    if len(mat) == 0 || r * c != len(mat) * len(mat[0]) {
        return mat
    }

    var row, col int
    result := make([][]int, r)
    result[0] = make([]int, c)
    for _, matRow := range mat {
        for _, num := range matRow {
            if col == c {
                col = 0
                row++
                result[row] = make([]int, c)
            }

            result[row][col] = num
            col++
        }
    }

    return result
}