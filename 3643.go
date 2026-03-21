func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
    for i := range k / 2 {
        for j := range k {
            grid[x+i][y+j], grid[x+k-i-1][y+j] = grid[x+k-i-1][y+j], grid[x+i][y+j]
        }
    }
    return grid
}
