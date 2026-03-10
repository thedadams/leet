type MyQueue struct {
    p, q []int
}


func Constructor() MyQueue {
    return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
    this.p = append([]int{x}, this.p...)
}


func (this *MyQueue) Pop() int {
    if len(this.q) == 0 {
        for i := range len(this.p) {
            this.q = append([]int{this.p[i]}, this.q...)
        }

        this.p = nil
    }

    num := this.q[0]
    this.q = this.q[1:]
    return num
}


func (this *MyQueue) Peek() int {
    if len(this.q) == 0 {
        for i := range len(this.p) {
            this.q = append([]int{this.p[i]}, this.q...)
        }

        this.p = nil
    }

    num := this.q[0]
    return num
}


func (this *MyQueue) Empty() bool {
    return len(this.p) == 0 && len(this.q) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */