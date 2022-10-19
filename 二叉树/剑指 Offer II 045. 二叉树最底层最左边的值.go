package 二叉树

func findBottomLeftValue(root *TreeNode) int {
	q := []*TreeNode{}
	q = append(q, root)
	ret := root.Val
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			top := q[0]
			q = q[1:]
			if i == 0 {
				ret = top.Val
			}

			if top.Left != nil {
				q = append(q, top.Left)
			}

			if top.Right != nil {
				q = append(q, top.Right)
			}
		}

	}

	return ret
}
