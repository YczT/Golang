package main

type StackOfPlates struct {
	cap    int
	stacks [][]int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{
		cap:    cap,
		stacks: make([][]int, 0),
	}
}

func (this *StackOfPlates) Push(val int) {
	if this.cap == 0 { // 栈长度为0，终止
		return
	}
	stack := make([]int, 0)
	if len(this.stacks) == 0 { // 创建第一个栈
		stack = append(stack, val)
		this.stacks = append(this.stacks, stack)
	} else {
		stack = this.stacks[len(this.stacks)-1]
		if len(stack) < this.cap {
			stack = append(stack, val)
			this.stacks[len(this.stacks)-1] = stack
		} else { // 栈满的时候创建新的栈
			stack = []int{val}
			this.stacks = append(this.stacks, stack)
		}
	}
}

func (this *StackOfPlates) Pop() int {
	if len(this.stacks) == 0 {
		return -1
	}

	stack := this.stacks[len(this.stacks)-1]
	top := -1
	if len(stack) > 0 {
		top = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
	}

	if len(stack) == 0 { // 栈空了要删掉
		this.stacks = this.stacks[0 : len(this.stacks)-1]
	} else {
		this.stacks[len(this.stacks)-1] = stack
	}

	return top
}

func (this *StackOfPlates) PopAt(index int) int {
	// 指定栈下标不得大于等于总栈数
	if index >= len(this.stacks) {
		return -1
	}

	stack := this.stacks[index]
	top := -1
	if len(stack) > 0 {
		top = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
	}

	if len(stack) == 0 { // 栈空要删除
		if index == len(this.stacks)-1 { // 若指定的栈是最后一个栈，直接删除
			this.stacks = this.stacks[0:index]
		} else { // 不是最后一个栈，使用切片特性，组成新的总栈
			this.stacks = append(this.stacks[0:index], this.stacks[index+1:]...)
		}
	} else {
		this.stacks[index] = stack
	}
	return top
}
