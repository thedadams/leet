func findClosest(x, y, z int) int {
    diff := int64(math.Abs(float64(x - z))) - int64(math.Abs(float64(y - z)))
    if diff > 0 {
        return 2
    } else if diff < 0 {
        return 1
    }
    return 0
}