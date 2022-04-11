package main

/*
给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的
深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。

示例：
输入：[1,2,3,4,5,null,7,8]

        1
       /  \
      2    3
     / \    \
    4   5    7
   /
  8

输出：[[1],[2,3],[4,5,7],[8]]
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type TreeNode43 struct {
	Val   int
	Left  *TreeNode43
	Right *TreeNode43
}
type ListNode43 struct {
	Val  int
	Next *ListNode43
}

//func getDeepthFunc(tree *TreeNode43) int {
//	deep, ans := 0, 0
//	if tree == nil {
//		return 0
//	}
//	var f func(node *TreeNode43)
//	f = func(node *TreeNode43) {
//		if node == nil {
//			return
//		}
//		deep++
//		f(node.Left)
//		f(node.Right)
//		ans = max(deep, ans)
//		deep--
//	}
//	f(tree)
//	return ans
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
//func exponent(a, n uint64) uint64 {
//	result := uint64(1)
//	for i := n; i > 0; i >>= 1 {
//		if i&1 != 0 {
//			result *= a
//		}
//		a *= a
//	}
//	return result
//}
//
//var i = 0

func listOfDepth(tree *TreeNode43) []*ListNode43 {
	if tree == nil {
		return []*ListNode43{}
	}
	ans := []*ListNode43{}
	queue := []*TreeNode43{}
	queue = append(queue, tree)

	for len(queue) > 0 {
		curSize := len(queue)        //当前层的尺寸
		nextLevel := []*TreeNode43{} //存储下一层节点

		head := new(ListNode43)
		curNode := head
		//遍历当前层
		for i := 0; i < curSize; i++ {
			curNode.Next = &ListNode43{Val: queue[i].Val}
			curNode = curNode.Next //后移
			//左右节点入队
			if queue[i].Left != nil {
				nextLevel = append(nextLevel, queue[i].Left)
			}
			if queue[i].Right != nil {
				nextLevel = append(nextLevel, queue[i].Right)
			}
		}
		queue = nextLevel
		ans = append(ans, head.Next)
	}
	return ans
}

func main() {

}
