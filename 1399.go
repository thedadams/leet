func countLargestGroup(n int) int {
    groups := make(map[int]int)
    var sum, largestSize int
    
    for i := 1; i <= n; i++ {
        sum = sumDigits(i)
        groups[sum]++
        largestSize = max(largestSize, groups[sum])
    }

    sum = 0
    for _, i := range groups {
        if i == largestSize {
            sum++
        }
    }

    return sum
}

func sumDigits(n int) int {
    var sum int
    for n > 0 {
        sum += n % 10
        n /= 10
    }

    return sum
}