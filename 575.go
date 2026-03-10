func distributeCandies(candyType []int) int {
    canEat := len(candyType) / 2
    seen := make(map[int]struct{}, len(candyType))

    var count int
    for _, t := range candyType {
        if _, ok := seen[t]; ok {
            continue
        }

        count++
        seen[t] = struct{}{}
        if canEat == count {
            return count
        }
    }

    return count
}