package 二叉树

import "math"

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{}
	q = append(q, root)
	ret := []int{}
	for len(q) > 0 {
		l := len(q)
		levelMax := math.MinInt32
		for i := 0; i < l; i++ {
			top := q[0]
			q = q[1:]

			if top.Val > levelMax {
				levelMax = top.Val
			}
			if top.Left != nil {
				q = append(q, top.Left)
			}

			if top.Right != nil {
				q = append(q, top.Right)
			}
		}

		ret = append(ret, levelMax)
	}

	return ret
}
