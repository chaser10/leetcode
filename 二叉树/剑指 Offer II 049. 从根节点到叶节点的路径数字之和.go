package 二叉树

func sumNumbers(root *TreeNode) int {
	sum := 0
	var dfs func(root *TreeNode, father int)
	dfs = func(root *TreeNode, father int) {
		if root == nil {
			return
		}

		cur := father*10 + root.Val
		if root.Left == nil && root.Right == nil {
			sum += cur
		}
		dfs(root.Left, cur)
		dfs(root.Right, cur)
	}

	dfs(root, 0)
	return sum
}
