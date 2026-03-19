type KthLargest struct {
	nums []int
	k    int
}

func (this *KthLargest) Len() int {
	return len(this.nums)
}

func (this *KthLargest) Less(i, j int) bool {
	return this.nums[i] < this.nums[j]
}

func (this *KthLargest) Swap(i, j int) {
	this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
}

func (this *KthLargest) Push(x any) {
	this.nums = append(this.nums, x.(int))
}

func (this *KthLargest) Pop() any {
	p := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return p
}

func Constructor(k int, nums []int) KthLargest {
	if grow := k + 1 - len(nums); grow > 0 {
		nums = slices.Grow(nums, grow)
        for range grow {
            nums = append(nums, -10001)
        }
	}
	kth := KthLargest{
		nums: nums,
		k:    k,
	}

	heap.Init(&kth)
    for kth.Len() > k {
        heap.Pop(&kth)
    }

	return kth
}

func (this *KthLargest) Add(val int) int {
	if val >= this.nums[0] {
		heap.Push(this, val)
		if len(this.nums) > this.k {
			heap.Pop(this)
		}
	}

	return this.nums[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
