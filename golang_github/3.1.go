package main

/*
三合一。描述如何只用一个数组来实现三个栈。

你应该实现push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、
peek(stackNum)方法。stackNum表示栈下标，value表示压入的值。
构造函数会传入一个stackSize参数，代表每个栈的大小。

示例1:
 输入：
["TripleInOne", "push", "push", "pop", "pop", "pop", "isEmpty"]
[[1], [0, 1], [0, 2], [0], [0], [0], [0]]
 输出：
[null, null, null, 1, -1, -1, true]
说明：当栈为空时`pop, peek`返回-1，当栈满时`push`不压入元素。

示例2:
 输入：
["TripleInOne", "push", "push", "push", "pop", "pop", "pop", "peek"]
[[2], [0, 1], [0, 2], [0, 3], [0], [0], [0], [0]]
 输出：
[null, null, null, null, 2, 1, -1, -1]
*/

type TripleInOne struct {
	Stack    []int
	StackLen int
	s1       int
	s2       int
	s3       int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		Stack:    make([]int, stackSize*3),
		StackLen: stackSize,
		s1:       0,
		s2:       stackSize,
		s3:       stackSize * 2,
	}
}

func (this *TripleInOne) Push(stackNum int, value int) { //stackNum决定传入哪一个栈
	if stackNum == 0 {
		if this.s1 == this.StackLen {
			return //栈满
		}
		this.Stack[this.s1] = value
		this.s1++
	} else if stackNum == 1 {
		if this.s2 == this.StackLen*2 {
			return
		}
		this.Stack[this.s2] = value
		this.s2++
	} else {
		if this.s3 == this.StackLen*3 {
			return
		}
		this.Stack[this.s3] = value
		this.s3++
	}
}

func (this *TripleInOne) Pop(stackNum int) int {
	if this.IsEmpty(stackNum) {
		return -1
	}
	if stackNum == 0 {
		tempval := this.Stack[this.s1-1]
		this.s1--
		return tempval
	} else if stackNum == 1 {
		tempval := this.Stack[this.s2-1]
		this.s2--
		return tempval
	} else {
		tempval := this.Stack[this.s3-1]
		this.s3--
		return tempval
	}

}

func (this *TripleInOne) Peek(stackNum int) int { //peek 不改变栈的值(不删除栈顶的值)，pop会把栈顶的值删除
	if this.IsEmpty(stackNum) {
		return -1
	}
	if stackNum == 0 {
		return this.Stack[this.s1-1]
	} else if stackNum == 1 {
		return this.Stack[this.s2-1]
	} else {
		return this.Stack[this.s3-1]
	}
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	if stackNum == 0 {
		return this.s1 == 0
	} else if stackNum == 1 {
		return this.s2 == this.StackLen
	} else {
		return this.s3 == this.StackLen*2
	}
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
