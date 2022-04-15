package main

import "math"

/*
实现一个函数，检查一棵二叉树是否为二叉搜索树。


示例1:
输入:
    2
   / \
  1   3
输出: true

示例2:
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false

解释: 输入为: [5,1,4,null,null,3,6]。
   根节点的值为 5 ，但是其右子节点值为 4 。

*/

type TreeNode5 struct {
	Val   int
	Left  *TreeNode5
	Right *TreeNode5
}

func isValidBST(root *TreeNode5) bool {
	pre = math.MaxInt64
	return DFS_5(root)

}

var pre int

func DFS_5(root *TreeNode5) bool {
	if root == nil {
		return true
	}
	left := DFS_5(root.Left)
	if pre != math.MaxInt64 && root.Val <= pre {
		return false
	}
	pre = root.Val
	right := DFS_5(root.Right)
	return left && right
}

func main() {

}
