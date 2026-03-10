type MyStack struct {
    p, q []int
}


func Constructor() MyStack {
    return MyStack{}
}


func (this *MyStack) Push(x int)  {
    this.p = append(this.p, x)
}


func (this *MyStack) Pop() int {
    for i := range len(this.p) - 1{
        this.q = append(this.q, this.p[i])
    }

    num := this.p[len(this.p) - 1]

    this.p = this.q
    this.q = nil
    return num
}


func (this *MyStack) Top() int {
    num := this.Pop()
    this.Push(num)
    return num
}


func (this *MyStack) Empty() bool {
    return len(this.p) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */