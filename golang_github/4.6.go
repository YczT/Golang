package main

/*
设计一个算法，找出二叉搜索树(二叉排序树)中指定节点的“下一个”节点（也即中序后继）。
如果指定节点没有对应的“下一个”节点，则返回null。

示例 1:
输入: root = [2,1,3], p = 1

  2
 / \
1   3

输出: 2

示例 2:
输入: root = [5,3,6,2,4,null,null,1], p = 6

      5
     / \
    3   6
   / \
  2   4
 /
1

输出: null

*/

type TreeNode6 struct {
	Val   int
	Left  *TreeNode6
	Right *TreeNode6
}

func inorderSuccessor(root *TreeNode6, p *TreeNode6) *TreeNode6 {
	if root == nil {
		return nil
	}
	var storey *TreeNode6
	for root != nil {
		if root.Val > p.Val {
			storey = root
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return storey
}

func main() {

}
