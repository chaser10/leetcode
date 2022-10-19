package 二叉树

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	tree []*TreeNode
	root *TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	tree := []*TreeNode{}
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		top := q[0]
		tree = append(tree, top)
		q = q[1:]
		if top.Left != nil {
			q = append(q, top.Left)
		}

		if top.Right != nil {
			q = append(q, top.Right)
		}
	}

	fmt.Println(tree)
	return CBTInserter{tree, root}
}

func (this *CBTInserter) Insert(v int) int {
	node := &TreeNode{v, nil, nil}
	father := this.tree[(len(this.tree)+1)/2-1]
	if father.Left == nil {
		father.Left = node
	} else {
		father.Right = node
	}

	this.tree = append(this.tree, node)

	return father.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
