package 链表

//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	p := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//	return p
//}

func reverseList(head *ListNode) *ListNode {
	pre := &ListNode{}
	for cur := head; cur != nil; cur = cur.Next {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
