package main

import "fmt"

/*
给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的
二叉搜索树。

示例:
给定有序数组: [-10,-3,0,5,9],
一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

          0
         / \
       -3   9
       /   /
     -10  5
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

func main() {
	num := []int{-10, -3, 0, 5, 9}
	ans := sortedArrayToBST(num)
	fmt.Printf("%v\n", ans)
}
