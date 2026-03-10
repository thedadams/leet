func mySqrt(x int) int {
    var i int
    for {
        if i*i > x {
            return i - 1
        }
        i++
    }
}