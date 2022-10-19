package 二叉树

func convertBST(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode)
	cur := 0
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Right)
		cur += root.Val
		root.Val = cur
		dfs(root.Left)
		return
	}

	dfs(root)

	return root
}
