package main

/*
实现一个函数，检查二叉树是否平衡。在这个问题中，平衡树的定义
如下：任意一个节点，其两棵子树的高度差不超过 1。

示例 1:
给定二叉树 [3,9,20,null,null,15,7]
    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:
给定二叉树 [1,2,2,3,3,null,null,4,4]
      1
     / \
    2   2
   / \
  3   3
 / \
4   4
返回false 。
*/

type TreeNode4 struct {
	Val   int
	Left  *TreeNode4
	Right *TreeNode4
}

func isBalanced(root *TreeNode4) bool {
	temp := root
	bl, _ := DFS_4(temp)
	return bl
}

func DFS_4(root *TreeNode4) (bool, int) {
	if root == nil {
		return true, 0
	}
	leftdeep, lefth := DFS_4(root.Left)
	if !leftdeep {
		return false, 0
	}
	rightdeep, righth := DFS_4(root.Right)
	if !rightdeep {
		return false, 0
	}
	if balance(lefth, righth) > 1 {
		return false, 0
	}
	return true, max(lefth, righth) + 1
}
func balance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
