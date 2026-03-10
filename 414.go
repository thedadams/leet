func thirdMax(nums []int) int {
    maxes := make([]int, 0)

    var idx, maxIdx int
    for i := range 3 {
        idx, maxIdx = -1, -1

        for j := range nums {
            if (maxIdx == -1 || nums[j] > nums[maxIdx]) && (i == 0 || nums[j] < nums[maxes[i-1]]) {
                idx = j
                maxIdx = j
            }
        }
        
        if idx == -1 {
            return nums[maxes[0]]
        }

        maxes = append(maxes, idx)
    }

    return nums[maxes[2]]
}