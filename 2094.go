func findEvenNumbers(digits []int) []int {
    digitCounts := make(map[int]int, len(digits))
    for _, i := range digits {
        digitCounts[i]++
    }

    var (
        result []int
        ones, tens, hundreds int
        onesCount, tensCount, hundredsCount int
    )
    for i := 100; i < 1000; i += 2 {
        ones, tens, hundreds = i % 10, (i / 10) % 10, (i / 100) % 10
        onesCount, tensCount, hundredsCount = digitCounts[ones], digitCounts[tens], digitCounts[hundreds]

        if onesCount*tensCount*hundredsCount == 0 {
            continue
        }

        if ones == tens && onesCount == 1 ||
        ones == hundreds && onesCount == 1 ||
        tens == hundreds && tensCount == 1 ||
        ones == tens && tens == hundreds && onesCount == 2 {
            continue
        }

        result = append(result, i)
    }

    return result
}