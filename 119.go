func getRow(rowIndex int) []int {
    row := make([]int, rowIndex + 1)
    
    for i := range rowIndex / 2 + 1 {
        row[i] = halfChoose(rowIndex, i)
        row[rowIndex - i] = row[i]
    }

    return row
}

func halfChoose(n, k int) int {
    var fact int64 = 1
    for i := n; i > n - k; i-- {
        fact *= int64(i)
        fact /= int64(n - i + 1)
    }

    return int(fact)
}