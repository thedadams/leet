func findMedianSortedArrays(nums1, nums2 []int) float64 {
    total := len(nums1) + len(nums2)
    midIdx := total / 2
    var mid1, mid2 int

    for i, j := 0, 0; i + j <= midIdx; {
        mid1 = mid2
        if i < len(nums1) {
            mid2 = nums1[i]
        } else {
            mid2 = nums2[j]
            j++
            continue
        }

        if j >= len(nums2) {
            i++
            continue
        }

        if mid2 > nums2[j] {
            mid2 = nums2[j]
            j++
        } else {
            i++
        }
    }

    if (len(nums1) + len(nums2)) % 2 == 0 {
        return float64(mid1 + mid2) / 2.0
    }

    return float64(mid2)
}