type NumArray struct {
    sums []int
}


func Constructor(nums []int) NumArray {
    sums := make([]int, len(nums) + 1)
    for i := 1; i < len(sums); i++ {
        sums[i] = sums[i-1] + nums[i-1]
    }

    return NumArray{sums}
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.sums[right+1] - this.sums[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */