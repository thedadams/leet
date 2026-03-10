func nextGreaterElement(nums1, nums2 []int) []int {
    stack := make([]int, 0, len(nums2))
    next := make(map[int]int, len(nums2))
    for _, i := range nums2 {
        for len(stack) > 0 && i > stack[len(stack) - 1] {
            next[stack[len(stack) - 1]] = i
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, i)
    }

    for _, i := range stack {
        next[i] = -1
    }

    for i := range nums1 {
        nums1[i] = next[nums1[i]]
    }
    
    return nums1
}