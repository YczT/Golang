package main

/*
检查子树。你有两棵非常大的二叉树：T1，有几万个节点；T2，有几万个节点。设计一个算法，判断 T2 是否为 T1 的子树。
如果 T1 有这么一个节点 n，其子树与 T2 一模一样，则 T2 为 T1 的子树，也就是说，从节点 n 处把树砍断，得到的树
与 T2 完全相同。
注意：此题相对书上原题略有改动。

示例1:
 输入：t1 = [1, 2, 3], t2 = [2]
 输出：true

示例2:
 输入：t1 = [1, null, 2, 4], t2 = [3, 2]
 输出：false
*/

type TreeNode10 struct {
	Val   int
	Left  *TreeNode10
	Right *TreeNode10
}

func checkSubTree(t1 *TreeNode10, t2 *TreeNode10) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil && t2 != nil {
		return false
	}
	if t1 != nil && t2 == nil {
		return false
	}
	if t1.Val == t2.Val {
		return checkSubTree(t1.Left, t2.Left) && checkSubTree(t1.Right, t2.Right)
	} else {
		return checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
	}
}

func DFS10val(t1 *TreeNode10, t2 *TreeNode10) {

}
func main() {

}
