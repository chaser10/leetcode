package 二叉树

func findTarget(root *TreeNode, k int) bool {
	var dfs func(root *TreeNode)
	m := map[int]int{}
	ret := false
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if _, ok := m[k-root.Val]; ok {
			ret = true
		}

		m[root.Val] = 1
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)

	return ret
}
