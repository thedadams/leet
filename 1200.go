func minimumAbsDifference(arr []int) [][]int {
    originalLen := len(arr)
    arr, m := countingSort(arr)
    var (
        result [][]int
        diff int
        minDiff = math.MaxInt
        found int
    )
    for i := m+1000000; i < len(arr) && found < originalLen; {
        for arr[i] == 0 {
            i++
            if i == len(arr) - 1 {
                return result
            }
        }
        found++
        if found == originalLen {
            return result
        }

        j := i
        if arr[i] == 1 {
            j++
        }
        for arr[j] == 0 {
            j++
            if j == len(arr) {
                return result
            }
        }

        diff = j - i
        if diff < minDiff {
            minDiff = diff
            result = [][]int{{i-1000000, j-1000000}}
        } else if diff == minDiff {
            result = append(result, []int{i-1000000, j-1000000})
        }

        i = j
    }

    return result
}

func countingSort(arr []int) ([]int, int) {
    result := make([]int, 2000001)
    m := arr[0]
    for _, num := range arr {
        result[num+1000000]++
        m = min(m, num)   
    }

    return result, m
}