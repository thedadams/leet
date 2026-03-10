func intersection(nums1 []int, nums2 []int) []int {
    intersection := make(map[int]struct{}, len(nums1))
    for _, num := range nums1 {
        intersection[num] = struct{}{}
    }

    var ok bool
    nums1 = nil
    for _, num := range nums2 {
        if _, ok = intersection[num]; ok {
            delete(intersection, num)
            nums1 = append(nums1, num)
        }
    }

    return nums1
}