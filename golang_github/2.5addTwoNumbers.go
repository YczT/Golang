package main

import "fmt"

/*
给定两个用链表表示的整数，每个节点包含一个数位。
这些数位是反向存放的，也就是个位排在链表首部。
编写函数对这两个整数求和，并用链表形式返回结果。
*/

type ListNode4 struct {
	Val  int
	Next *ListNode4
}

func makeListNode4(nums []int) *ListNode4 {
	//创建链表
	if len(nums) == 0 {
		return nil
	}
	res := &ListNode4{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode4{Val: nums[i]}
		temp = temp.Next
	}
	return res
}
func addTwoNumbers(l1 *ListNode4, l2 *ListNode4) *ListNode4 {
	temp := &ListNode4{}
	lst := temp
	up := false
	for l1 != nil || l2 != nil || up {
		val := 0
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		if up {
			val++
		}
		up = val > 9
		node := &ListNode4{Val: val % 10}
		lst.Next = node //lst本身首值是零，要往前走一位
		lst = node
	}
	return temp.Next
}
func main() {
	list1 := makeListNode4([]int{1, 8, 5})
	list2 := makeListNode4([]int{2, 3, 5})
	ans := addTwoNumbers(list1, list2)
	for ans != nil {
		fmt.Printf("%v", ans.Val)
		ans = ans.Next
	}
}
