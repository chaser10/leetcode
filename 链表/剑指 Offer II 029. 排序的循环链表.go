package 链表

import "fmt"

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		node := &Node{Val: x}
		node.Next = node
		return node
	}
	cur := aNode
	tail := aNode
	for {
		if cur.Val > tail.Val {
			tail = cur
		}

		if cur == aNode {
			break
		}

		if x >= cur.Val && x <= cur.Next.Val {
			node := &Node{Val: x, Next: cur.Next}
			cur.Next = node
			fmt.Println(cur.Val)
			return aNode
		}

		cur = cur.Next
	}

	tail.Next = &Node{Val: x, Next: tail.Next}
	return aNode
}
