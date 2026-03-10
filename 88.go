func merge(nums1 []int, m int, nums2 []int, n int)  {
   nums := make([]int , m)
   copy(nums, nums1) 

    var i, j, k int
   for ; i < m && j < n; k++ {
    if nums[i] <= nums2[j] {
        nums1[k] = nums[i]
        i++
    } else {
        nums1[k] = nums2[j]
        j++
    }
   }

   for i < m {
    nums1[k] = nums[i]
    i++
    k++
   }
   for j < n {
    nums1[k] = nums2[j]
    j++
    k++
   }
}