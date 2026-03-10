func findRelativeRanks(score []int) []string {
    ranks := make([]string, len(score))
    var m int
    for i := range len(score) {
        m = i
        for j := range score {
            if score[j] > score[m] {
                m = j
            }
        }
        switch i {
            case 0:
            ranks[m] = "Gold Medal"
            case 1:
            ranks[m] = "Silver Medal"
            case 2:
            ranks[m] = "Bronze Medal"
            default:
            ranks[m] = fmt.Sprintf("%d", i + 1)
        }

        score[m] = -1
    }

    return ranks
}