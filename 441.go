func arrangeCoins(n int) int {
    return (-1 + int(math.Sqrt(float64(1 + 8 * n)))) / 2
}