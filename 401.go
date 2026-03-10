func readBinaryWatch(turnedOn int) []string {
    var result []string

    var hourBits int
    for i := range 12 {
        hourBits = countBits(i)
        if hourBits > turnedOn {
            continue
        }

        for j := range 60 {
            if hourBits + countBits(j) == turnedOn {
                result = append(result, fmt.Sprintf("%d:%02d", i, j))
            }
        }
    }
    
    return result
}

func countBits(num int) int {
    var ans int
    for num > 0 {
        ans += num & 1
        num >>= 1
    }

    return ans
}