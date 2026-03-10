func maximumProduct(nums []int) int {
    maximums := make(map[int]struct{}, 3)
    minimums := make(map[int]struct{}, 3)
    ma, mi := 1, 1
    for j := range 3 {
        maxIdx, minIdx := -1, -1
        for idx, i := range nums {
            if _, ok := maximums[idx]; !ok && (maxIdx == -1 || nums[maxIdx] < i) {
                maxIdx = idx
            }
            if _, ok := minimums[idx]; !ok && (minIdx == -1 || nums[minIdx] > i) {
                minIdx = idx
            }
        }

        maximums[maxIdx] = struct{}{}
        minimums[minIdx] = struct{}{}
        ma = nums[maxIdx]
        if j < 2 {
            mi *= nums[minIdx]
        }
    }

    r := 1
    for i := range maximums {
        n := nums[i]
        r *= n
        ma = max(ma, n)
    }

    result := 1
    for i := range minimums {
        result *= nums[i]
    }

    if mi == 0 {
        return max(result, r)
    }

    return max(result, r, mi * ma)
}