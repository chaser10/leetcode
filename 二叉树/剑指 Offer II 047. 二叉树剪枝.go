package 二叉树

import "fmt"

func pruneTree(root *TreeNode) *TreeNode {
	var isZeroTree func(root *TreeNode) bool
	isZeroTree = func(root *TreeNode) bool {
		if root == nil || root.Val == 0 && root.Left == nil && root.Right == nil {
			fmt.Println(root.Val)
			return true
		}

		left := isZeroTree(root.Left)
		right := isZeroTree(root.Right)
		if left {
			root.Left = nil
		}

		if right {
			root.Right = nil
		}
		return left && right && root.Val == 0
	}

	isZeroTree(root)

	return root
}
