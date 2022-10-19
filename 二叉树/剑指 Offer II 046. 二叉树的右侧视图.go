package 二叉树

func rightSideView(root *TreeNode) []int {
	curHeight := -1
	ret := []int{}
	var dfs func(root *TreeNode, height int)
	dfs = func(root *TreeNode, height int) {
		if root == nil {
			return
		}
		if height > curHeight {
			ret = append(ret, root.Val)
			curHeight = height
		}

		dfs(root.Right, height+1)
		dfs(root.Left, height+1)
		return
	}

	dfs(root, 0)
	return ret
}
