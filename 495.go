func findPoisonedDuration(timeSeries []int, duration int) int {
    var total int
    if len(timeSeries) == 0 {
        return total
    }

    // This takes care of the last element in the list. Since it last,
    // the poison will last for the full duration.
    total += duration

    for i := 1; i < len(timeSeries); i++ {
        if timeSeries[i] - timeSeries[i - 1] < duration {
            total += timeSeries[i] - timeSeries[i - 1]
        } else {
            total += duration
        }
    }
    
    return total
}