package main

import "fmt"

/*
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	temp := &ListNode{Next: head}
	maplistnode := make(map[int]int)
	for i := temp; i != nil && i.Next != nil; {
		nextval := i.Next.Val
		if _, e := maplistnode[nextval]; e {
			i.Next = i.Next.Next
		} else {
			maplistnode[nextval] = 1
			i = i.Next
		}
	}
	return temp.Next
}
func makeListNode(nums []int) *ListNode {

	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}

	temp := res

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val: nums[i]}
		temp = temp.Next
	}

	return res
}
func main() {
	list := makeListNode([]int{1, 4, 5, 2, 6, 3, 2, 1, 5, 5, 5, 6})
	ans := removeDuplicateNodes(list)
	for ans != nil {
		fmt.Printf("%v", ans.Val)
		ans = ans.Next
	}
}
