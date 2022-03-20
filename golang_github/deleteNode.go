package main

import "fmt"

/*
若链表中的某个节点，既不是链表头节点，也不是链表尾节点，则称其为该链表的「中间节点」。
假定已知链表的某一个中间节点，请实现一种算法，将该节点从链表中删除。
例如，传入节点c（位于单向链表a->b->c->d->e->f中），将其删除后，剩余链表为a->b->d->e->f
*/

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func deleteNode(node *ListNode2) *ListNode2 {
	//for {
	//	node.Val = node.Next.Val
	//	if node.Next.Next == nil {
	//		node.Next = nil
	//		break
	//	} else {
	//		node = node.Next
	//	}
	//}
	*node = *node.Next

	return node
}
func main() {
	list := makeListNode2([]int{1, 2, 9, 4, 5, 2, 6})
	ans := deleteNode(list)
	for ans != nil {
		fmt.Printf("%v", ans.Val)
		ans = ans.Next
	}
}
func makeListNode2(nums []int) *ListNode2 {
	//创建链表
	if len(nums) == 0 {
		return nil
	}
	res := &ListNode2{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode2{Val: nums[i]}
		temp = temp.Next
	}
	return res
}
