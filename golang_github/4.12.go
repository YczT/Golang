package main

/*
给定一棵二叉树，其中每个节点都含有一个整数数值(该值或正或负)。设计一个算法，打印节点数值总和等于某个给定值的所有路
径的数量。注意，路径不一定非得从二叉树的根节点或叶节点开始或结束，但是其方向必须向下(只能从父节点指向子节点方向)。

示例:
给定如下二叉树，以及目标和sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:
3
解释：和为 22的路径有：[5,4,11,2], [5,8,4,5], [4,11,7]

*/

type TreeNode12 struct {
	Val   int
	Left  *TreeNode12
	Right *TreeNode12
}

func pathSum(root *TreeNode12, sum int) int {
	count := 0 //记录值相等的次数，返回一次+1
	if root == nil {
		return 0
	}
	count = DFS12(root, sum)
	count += pathSum(root.Left, sum)  //以当前根节点的左孩子结点为根节点继续遍历
	count += pathSum(root.Right, sum) //以当前根节点的右孩子结点为根节点继续遍历
	return count

}
func DFS12(root *TreeNode12, sum int) (ans int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == sum {
		ans++
	}
	ans += DFS12(root.Left, sum-val)
	ans += DFS12(root.Right, sum-val)
	return
}
func main() {

}
