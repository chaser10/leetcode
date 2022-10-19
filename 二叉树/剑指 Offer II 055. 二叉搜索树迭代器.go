package 二叉树

type BSTIterator struct {
	st []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	st := []*TreeNode{}
	for root != nil {
		st = append(st, root)
		root = root.Left
	}

	return BSTIterator{st: st}
}

func (this *BSTIterator) Next() int {
	cur := this.st[len(this.st)-1]
	this.st = this.st[:len(this.st)-1]
	right := cur.Right
	for right != nil {
		this.st = append(this.st, right)
		right = right.Left
	}
	return cur.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.st) > 0
}
