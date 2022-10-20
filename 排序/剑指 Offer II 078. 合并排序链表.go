func mergeKLists(lists []*ListNode) *ListNode {
    return merge(lists, 0, len(lists) - 1)
}

func merge(lists []*ListNode, left int, right int) *ListNode {
    if left == right {
        return lists[left]
    }
    
    mid := (left + right) >> 1
    l1 := merge(lists, left, mid)
    l2 := merge(lists, mid + 1, right)

    dummy := &ListNode{}
    cur := dummy
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            cur.Next = l1
            l1 = l1.Next
        } else {
            cur.Next = l2
            l2 = l2.Next
        }

        cur = cur.Next
    }

    if l1 != nil {
        cur.Next = l1
    }

    if l2 != nil {
        cur.Next = l2
    }

    return dummy.Next
}