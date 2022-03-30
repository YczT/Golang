package main

import "fmt"

/*
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
*/

type ListNode6 struct {
	Val  int
	Next *ListNode6
}

func makeListNode6(nums []int) *ListNode6 {
	//创建链表
	if len(nums) == 0 {
		return nil
	}
	res := &ListNode6{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode6{Val: nums[i]}
		temp = temp.Next
	}
	return res
}
func getIntersectionNode(headA, headB *ListNode6) *ListNode6 {
	fast, slow := headA, headB
	temp1 := ListNode6{
		Next: headA,
	}
	temp2 := ListNode6{
		Next: headB,
	}
	lengthA, lengthB := 0, 0
	for temp1.Next != nil {
		lengthA++
		temp1.Next = temp1.Next.Next
	}
	for temp2.Next != nil {
		lengthB++
		temp2.Next = temp2.Next.Next
	}
	var step int
	if lengthA > lengthB {
		step = lengthA - lengthB
		fast, slow = headA, headB
	} else {
		step = lengthB - lengthA
		fast, slow = headB, headA
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
func main() {
	list1 := makeListNode6([]int{4, 1, 8, 4, 5})
	list2 := makeListNode6([]int{5, 0, 1, 8, 4, 5})
	ans := getIntersectionNode(list1, list2)
	fmt.Println(ans)
}
