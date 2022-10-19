package 二叉树

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	q := []*TreeNode{}
	q = append(q, root)
	sb := &strings.Builder{}
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		if top == nil {
			sb.WriteString("nil,")
			continue
		} else {
			sb.WriteString(strconv.Itoa(top.Val))
			sb.WriteByte(',')
		}
		q = append(q, top.Left)
		q = append(q, top.Right)

	}

	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	treeStrs := strings.Split(data, ",")
	if len(treeStrs) == 0 || treeStrs[0] == "nil" {
		return nil
	}

	rootVal, _ := strconv.Atoi(treeStrs[0])
	root := &TreeNode{rootVal, nil, nil}
	q := []*TreeNode{}
	q = append(q, root)
	treeStrs = treeStrs[1:]
	for len(q) > 0 {
		left := treeStrs[0]
		right := treeStrs[1]
		treeStrs = treeStrs[2:]
		top := q[0]
		q = q[1:]
		if left != "nil" {
			leftVal, _ := strconv.Atoi(left)
			leftNode := &TreeNode{leftVal, nil, nil}
			top.Left = leftNode
			q = append(q, leftNode)
		}

		if right != "nil" {
			rightVal, _ := strconv.Atoi(right)
			rightNode := &TreeNode{rightVal, nil, nil}
			top.Right = rightNode
			q = append(q, rightNode)
		}
	}

	return root
}
