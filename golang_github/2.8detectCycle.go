package main

import "fmt"

/*
给定一个链表，如果它是有环链表，实现一个算法返回环路的开头节点。若环不存在，请返回 null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中
的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则
在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

*/

type ListNode7 struct {
	Val  int
	Next *ListNode7
}

func makeListNode7(nums []int) *ListNode7 {
	//创建链表
	if len(nums) == 0 {
		return nil
	}
	res := &ListNode7{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode7{Val: nums[i]}
		temp = temp.Next
	}
	return res
}
func detectCycle(head *ListNode7) *ListNode7 {
	maphead := map[*ListNode7]struct{}{}
	for head != nil {
		if _, ok := maphead[head]; ok {
			return head
		}
		maphead[head] = struct{}{} //相当于把值初始化了
		head = head.Next
	}
	return nil
}
func main() {
	list := makeListNode7([]int{4, 1, 8, 4, 5})
	ans := detectCycle(list)
	fmt.Println(ans)
}
