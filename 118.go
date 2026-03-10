func generate(numRows int) [][]int {
    rows := [][]int{{1}}
   if numRows == 1 {
    return rows
   }

   for i := range numRows-1 {
    rows = append(rows, make([]int, i + 2))
    rows[i+1][0] = 1
    rows[i+1][i+1] = 1

    for j := 1; j < i+1; j++ {
        rows[i+1][j] = rows[i][j-1] + rows[i][j]
    }
   }

   return rows
}