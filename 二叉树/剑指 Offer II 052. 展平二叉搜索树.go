package 二叉树

func increasingBST(root *TreeNode) *TreeNode {
	treeSlice := []*TreeNode{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		treeSlice = append(treeSlice, root)
		dfs(root.Right)
	}

	dfs(root)
	for i := 1; i < len(treeSlice); i++ {
		preNode := treeSlice[i-1]
		preNode.Left = nil
		preNode.Right = treeSlice[i]
		if i == len(treeSlice)-1 {
			treeSlice[i].Left = nil
			treeSlice[i].Right = nil
		}
	}

	return treeSlice[0]
}
