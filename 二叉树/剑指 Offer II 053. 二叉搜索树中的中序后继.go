package 二叉树

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) *TreeNode
	var pre *TreeNode
	dfs = func(cur *TreeNode) *TreeNode {
		if cur == nil {
			return nil
		}

		left := dfs(cur.Left)
		if left != nil {
			return left
		}
		if pre != nil && pre.Val == p.Val {
			return cur
		}
		pre = cur
		return dfs(cur.Right)
	}

	return dfs(root)
}
