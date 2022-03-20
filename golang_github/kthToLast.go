package main

import "fmt"

/*
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
*/
func makeListNode1(nums []int) *ListNode1 {

	if len(nums) == 0 {
		return nil
	}

	res := &ListNode1{
		Val: nums[0],
	}

	temp := res

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode1{Val: nums[i]}
		temp = temp.Next
	}

	return res
}

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func kthToLast(head *ListNode1, k int) int {
	length := 0
	temp := &ListNode1{Next: head}
	for temp.Next != nil {
		length++
		temp.Next = temp.Next.Next
	}
	answer := length - k
	v := 0
	flag := 0
	for head != nil {
		if flag != answer {
			head = head.Next
			flag++

		} else {
			v = head.Val
			break
		}
	}
	return v
}
func main() {
	list := makeListNode1([]int{1, 2, 10, 4, 5})
	k := 3
	ans := kthToLast(list, k)
	fmt.Println(ans)
}
