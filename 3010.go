func minimumCost(nums []int) int {
    one, two := nums[1], nums[2]
    if one > two {
        one, two = two, one
    }
    
    for i := 3; i < len(nums); i++ {
        if nums[i] < one {
            two = one
            one = nums[i]
        } else if nums[i] < two {
            two = nums[i]
        }
    }

    return nums[0] + one + two
}