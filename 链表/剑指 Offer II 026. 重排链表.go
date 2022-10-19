package 链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	var reverseList func(head *ListNode) *ListNode
	reverseList = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}

		p := reverseList(head.Next)
		head.Next.Next = head
		head.Next = nil
		return p
	}

	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	l1 := head
	l2 := reverseList(slow.Next)
	slow.Next = nil
	for l1 != nil && l2 != nil {
		l1Next := l1.Next
		l1.Next = l2
		l2Next := l2.Next
		l2.Next = l1Next
		l1 = l1Next
		l2 = l2Next
	}

}
