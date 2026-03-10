func islandPerimeter(grid [][]int) int {
    var result int

    for i, row := range grid {
        for j, land := range row {
            if land == 0 {
                continue
            }

            result += 4
            if i > 0 && grid[i - 1][j] == 1 {
                result -= 2
            }
            if j > 0 && grid[i][j - 1] == 1 {
                result -= 2
            }
        }
    }

    return result
}