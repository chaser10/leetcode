package 链表

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	head := root

	for root != nil {
		next := root.Next
		if root.Child != nil {
			flattenChild := flatten(root.Child)
			root.Next = flattenChild
			flattenChild.Prev = root
			root.Child = nil
			tail := root
			for tail.Next != nil {
				tail = tail.Next
			}

			tail.Next = next
			if next != nil {
				next.Prev = tail
			}
		}

		root = next
	}

	return head
}
