package 二叉树

import "math"

func maxPathSum(root *TreeNode) int {
	ret := math.MinInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return math.MinInt32
		}

		left := dfs(root.Left)
		right := dfs(root.Right)
		childMax := max(left, right)
		curMax := max(root.Val, root.Val+childMax)
		ret = max(ret, childMax)
		ret = max(max(ret, curMax), root.Val+left+right)
		return curMax
	}

	dfs(root)
	return ret
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
