package main

/*
堆盘子。设想有一堆盘子，堆太高可能会倒下来。因此，在现实生活中，盘子堆到一定高度时，我
们就会另外堆一堆盘子。请实现数据结构SetOfStacks，模拟这种行为。SetOfStacks应该由多
个栈组成，并且在前一个栈填满时新建一个栈。此外，SetOfStacks.push()和SetOfStacks.pop()
应该与普通栈的操作方法相同（也就是说，pop()返回的值，应该跟只有一个栈时的情况一样）。
进阶：实现一个popAt(int index)方法，根据指定的子栈，执行pop操作。

当某个栈为空时，应当删除该栈。当栈中没有元素或不存在该栈时，pop，popAt应返回 -1.

示例1:
 输入：
["StackOfPlates", "push", "push", "popAt", "pop", "pop"]
[[1], [1], [2], [1], [], []]
 输出：
[null, null, null, 2, 1, -1]

示例2:
 输入：
["StackOfPlates", "push", "push", "push", "popAt", "popAt", "popAt"]
[[2], [1], [2], [3], [0], [0], [0]]
 输出：
[null, null, null, null, 2, 1, 3]
*/
type StackOfPlates struct {
	datastacks [][]int
	num        int
}

func Constructor2(cap int) StackOfPlates {
	return StackOfPlates{
		make([][]int, 0),
		cap,
	}
}

func (this *StackOfPlates) Push(val int) {
	if this.num == 0 {
		return
	}
	stack := make([]int, 0) //新建一个栈
	if len(this.datastacks) == 0 {
		stack = append(stack, val)                       //第一个小栈
		this.datastacks = append(this.datastacks, stack) //第一个大栈
	} else {
		stack = this.datastacks[len(this.datastacks)-1]
		if len(stack) < this.num { //没超过规定的小栈的大小，不用新建栈
			stack = append(stack, val)
			this.datastacks[len(this.datastacks)-1] = stack //把小栈拿出来赋值，再重新赋值给大栈
		} else {
			stack = []int{val} //超过规定大小，新建一个小栈，并赋值
			this.datastacks = append(this.datastacks, stack)
		}
	}
}

func (this *StackOfPlates) Pop() int {
	if len(this.datastacks) == 0 {
		return -1
	}
	stack := this.datastacks[len(this.datastacks)-1]
	top := -1
	if len(stack) > 0 { //小栈里面有数据，获取该数据的值，并删除该数据
		top = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
	}
	if len(stack) == 0 { //小栈里面没有数据，删掉小栈
		this.datastacks = this.datastacks[0 : len(this.datastacks)-1]
	} else {
		this.datastacks[len(this.datastacks)-1] = stack
	}
	return top
}

func (this *StackOfPlates) PopAt(index int) int {
	if index >= len(this.datastacks)-1 { //指定小栈的下标不得超过已有的
		return -1
	}
	stack := this.datastacks[index]
	top := -1
	if len(stack) > 0 { //小栈里面有数据，获取该数据的值，并删除该数据
		top = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
	}
	if len(stack) == 0 { //小栈里面没有数据，删掉小栈
		if index == len(this.datastacks)-1 { // 若指定的小栈是最后一个小栈，直接删除
			this.datastacks = this.datastacks[0:index]
		} else { // 不是最后一个小栈，使用切片特性，组成新的大栈
			this.datastacks = append(this.datastacks[0:index], this.datastacks[index+1:]...)
		}
	} else {
		this.datastacks[index] = stack
	}
	return top
}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
