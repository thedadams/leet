func findRotation(mat [][]int, target [][]int) bool {
    for range 4 {
        if equal(mat, target) {
            return true
        }
        rotate(mat)
    }
    return false
}

func equal(mat1, mat2 [][]int) bool {
    for i := range len(mat1) {
        for j := range len(mat1[i]) {
            if mat1[i][j] != mat2[i][j] {
                return false
            }
        }
    }

    return true
}

func rotate(mat [][]int) {
    var temp int
    n := len(mat)
    for i := range (n + 1) / 2 {
    	for j := range n / 2 {
			temp = mat[i][j]
			mat[i][j] = mat[n-1-j][i]
			mat[n-1-j][i] = mat[n-1-i][n-1-j]
			mat[n-1-i][n-1-j] = mat[j][n-1-i]
			mat[j][n-1-i] = temp
		}
    }
}
