package algorithms

//Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 继续送分，归并操作
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head, n *ListNode
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	n = head
	for {
		if l1 == nil {
			n.Next = l2
			break
		}
		if l2 == nil {
			n.Next = l1
			break
		}
		if l1.Val < l2.Val {
			n.Next = l1
			l1 = l1.Next
		} else {
			n.Next = l2
			l2 = l2.Next
		}
		n = n.Next
	}
	return head
}
