func intersect(nums1, nums2 []int) []int {
    intersection := make(map[int]int, len(nums1))
    for _, num := range nums1 {
        intersection[num]++
    }

    nums1 = nil
    for _, num := range nums2 {
        if intersection[num] > 0 {
            nums1 = append(nums1, num)
            intersection[num]--
        }
    }

    return nums1
}