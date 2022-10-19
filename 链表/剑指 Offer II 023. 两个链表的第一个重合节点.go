package 链表

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 == nil {
			l1 = headB
		} else {
			l1 = l1.Next
		}

		if l2 == nil {
			l2 = headA
		} else {
			l2 = l2.Next
		}
	}

	return l1
}
