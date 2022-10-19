package 链表

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	fmt.Println(length)
	cur := &ListNode{0, head}
	for i := 0; i < n-length; i++ {
		fmt.Println(cur.Val)
		cur = cur.Next
	}

	cur.Next = cur.Next.Next
	return head
}
