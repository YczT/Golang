package main

import "fmt"

/*
栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。最多只能使用一个其他的临时栈存放
数据，但不得将元素复制到别的数据结构（如数组）中。该栈支持如下操作：push、pop、peek
和 isEmpty。当栈为空时，peek返回 -1。

 输入：
["SortedStack", "push", "push", "peek", "pop", "peek"]
[[], [1], [2], [], [], []]
 输出：
[null,null,null,1,null,2]
*/

type SortedStack struct {
	Pristack []int
}

func Constructor5() SortedStack {
	return SortedStack{}
}

func (this *SortedStack) Push(val int) {
	length := len(this.Pristack)
	if length == 0 || this.Pristack[length-1] <= val {
		this.Pristack = append(this.Pristack, val)
		return
	}
	for length != 0 && val < this.Pristack[length-1] {
		length--
	}
	this.Pristack = append(this.Pristack, append([]int{val}, this.Pristack[length:]...)...)
}

func (this *SortedStack) Pop() {
	if this.IsEmpty() {
		return
	}
	this.Pristack = this.Pristack[1:]
}

func (this *SortedStack) Peek() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Pristack[0]
}

func (this *SortedStack) IsEmpty() bool {
	return len(this.Pristack) == 0
}

/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */
func main() {
	obj := Constructor5()
	obj.Push(1)
	obj.Push(2)
	param_1 := obj.Peek()
	param_2 := obj.IsEmpty()
	fmt.Printf("%v\n", param_1)
	fmt.Printf("%v\n", param_2)
}
