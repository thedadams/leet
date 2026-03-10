func findRestaurant(list1, list2 []string) []string {
    stringIndexes := make(map[string]int, len(list1) + len(list2))
    for i, s := range list1 {
        stringIndexes[s] = i
    }

    var (
        result []string
        minSum = math.MaxInt
    )
    for i, s := range list2 {
        idx, ok := stringIndexes[s]
        if ok {
            stringIndexes[s] = idx + i
            if idx := idx + i; idx == minSum {
                result = append(result, s)
            } else if idx < minSum {
                minSum = idx
                result = result[:0]
                result = append(result, s)
            }
        }
    }

    return result
}