package 链表

import "fmt"

func isPalindrome(head *ListNode) bool {
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

	l2 := reverseList(slow.Next)
	fmt.Println(l2)
	l1 := head
	for l2 != nil {
		if l1.Val != l2.Val {
			return false
		}

		l2 = l2.Next
		l1 = l1.Next
	}

	return true
}
