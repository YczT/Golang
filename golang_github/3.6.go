package main

import (
	"fmt"
)

/*
动物收容所。有家动物收容所只收容狗与猫，且严格遵守“先进先出”的原则。在收养
该收容所的动物时，收养人只能收养所有动物中“最老”（由其进入收容所的时间长短
而定）的动物，或者可以挑选猫或狗（同时必须收养此类动物中“最老”的）。换言之，
收养人不能自由挑选想收养的对象。请创建适用于这个系统的数据结构，实现各种操
作方法，比如enqueue、dequeueAny、dequeueDog和dequeueCat。
允许使用Java内置的LinkedList数据结构。
enqueue方法有一个animal参数，animal[0]代表动物编号，animal[1]代表动物
种类，其中 0 代表猫，1 代表狗。dequeue*方法返回一个列表[动物编号, 动物种
类]，若没有可以收养的动物，则返回[-1,-1]。

示例1:
输入：
["AnimalShelf", "enqueue", "enqueue", "dequeueCat", "dequeueDog", "dequeueAny"]
[[], [[0, 0]], [[1, 0]], [], [], []]
输出：
[null,null,null,[0,0],[-1,-1],[1,0]]

示例2:
输入：
["AnimalShelf", "enqueue", "enqueue", "enqueue", "dequeueDog", "dequeueCat", "dequeueAny"]
[[], [[0, 0]], [[1, 0]], [[2, 1]], [], [], []]
输出：
[null,null,null,null,[2,1],[0,0],[1,0]]
*/

type AnimalShelf struct {
	cats, dogs [][]int
}

func Constructor6() AnimalShelf {
	return AnimalShelf{}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	if animal[1] == 0 {
		this.cats = append(this.cats, animal)
	} else {
		this.dogs = append(this.dogs, animal)
	}
}

func (this *AnimalShelf) DequeueAny() []int {
	if len(this.cats) == 0 {
		return this.DequeueDog()
	}
	if len(this.dogs) == 0 {
		return this.DequeueCat()
	}
	if this.cats[0][0] < this.dogs[0][0] {
		return this.DequeueCat()
	}
	return this.DequeueDog()
}

func (this *AnimalShelf) DequeueDog() []int {
	if len(this.dogs) == 0 {
		return []int{-1, -1}
	}
	as := this.dogs[0]
	this.dogs = this.dogs[1:]
	return as
}

func (this *AnimalShelf) DequeueCat() []int {
	if len(this.cats) == 0 {
		return []int{-1, -1}
	}
	as := this.cats[0]
	this.cats = this.cats[1:]
	return as
}

/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */
func main() {
	obj := Constructor6()
	obj.Enqueue([]int{0, 0})
	obj.Enqueue([]int{1, 0})
	p1 := obj.DequeueCat()
	p2 := obj.DequeueDog()
	p3 := obj.DequeueAny()
	fmt.Printf("%v%v%v\n", p1, p2, p3)
}
