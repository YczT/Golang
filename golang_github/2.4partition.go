package main

import "fmt"

/*
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你不需要保留每个分区中各节点的初始相对位置。
例子1:
输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]
例子2：
输入：head = [2,1], x = 2
输出：[1,2]
*/

type ListNode3 struct {
	Val  int
	Next *ListNode3
}

func makeListNode3(nums []int) *ListNode3 {
	//创建链表
	if len(nums) == 0 {
		return nil
	}
	res := &ListNode3{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode3{Val: nums[i]}
		temp = temp.Next
	}
	return res
}

func partition(head *ListNode3, x int) *ListNode3 {
	if head == nil {
		return nil
	}
	temp := &ListNode3{Next: head}
	//aimpointer := temp.Next
	//for {
	//	if aimpointer != nil && aimpointer.Next != nil && aimpointer.Val != x {
	//		aimpointer = aimpointer.Next
	//	} else {
	//		aimpointer.Val = temp.Next.Val
	//		temp.Next.Val = x
	//		break
	//	}
	//}

	pointer := head
	for pointer != nil && pointer.Next != nil {
		if pointer.Next.Val < x {
			node := pointer.Next
			nxt := pointer.Next.Next
			node.Next = temp.Next
			temp.Next = node
			pointer.Next = nxt
		} else {
			pointer = pointer.Next
		}
	}
	return temp.Next
}
func main() {
	list := makeListNode3([]int{1, 4, 3, 2, 5, 2})
	param := 3
	ans := partition(list, param)
	for ans != nil {
		fmt.Printf("%v", ans.Val)
		ans = ans.Next
	}
}
