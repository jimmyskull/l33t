type MinStack struct {
    elements []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
    this.elements = append(this.elements, x)
}


func (this *MinStack) Pop()  {
    size := len(this.elements)
    if size > 0 {
        this.elements = this.elements[:size-1]
    }
}


func (this *MinStack) Top() int {
    return this.elements[len(this.elements)-1]
}


func (this *MinStack) GetMin() int {
    var min int
    for i, v := range this.elements {
        if i == 0 || v < min {
            min = v
        }
    }
    return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */