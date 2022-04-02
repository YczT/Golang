package main

/*
请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，该函数返回栈元素中的最小值。执行push、pop和min操作的时间复杂度必须为O(1)。

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
*/

type MinStack struct {
	Stackdata []int
	Stackmin  []int
	top       int
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{
		make([]int, 10010),
		make([]int, 10010),
		0,
	}
}
func min(v1 int, v2 int) int {
	if v1 > v2 {
		return v2
	} else {
		return v1
	}
}

func (this *MinStack) Push(x int) {
	this.top++
	if this.top == 1 {
		this.Stackdata[this.top] = x
		this.Stackmin[this.top] = x
	} else {
		this.Stackdata[this.top] = x
		this.Stackmin[this.top] = min(x, this.Stackmin[this.top-1])
	}
}

func (this *MinStack) Pop() {
	this.top--
}

func (this *MinStack) Top() int {
	return this.Stackdata[this.top]
}

func (this *MinStack) GetMin() int {
	return this.Stackmin[this.top]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
