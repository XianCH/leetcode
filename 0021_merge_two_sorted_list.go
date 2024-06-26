package leetcode

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoList(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoList(l1, l2.Next)
		return l2
	}
}

// func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
// 	head := &ListNode{}
// 	current := head
// 	for l1 != nil && l2 != nil {
// 		if l1.Val < l2.Val {
// 			current.Next = l1
// 			l1 = l1.Next
// 		} else {
// 			current.Next = l2
// 			l2 = l2.Next
// 		}
// 		current = current.Next
// 	}
// 	if l1 == nil {
// 		current.Next = l2
// 	} else {
// 		current.Next = l1
// 	}
// 	return head.Next
// }
