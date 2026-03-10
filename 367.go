func isPerfectSquare(num int) bool {
    start, end := 2, num
    var mid, square int
    for start != end {
        mid = (start + end) / 2
        square = mid * mid
        if square == num {
            return true
        }
        if square < num {
            start = mid + 1
        } else {
            end = mid
        }
    }

    return false
}