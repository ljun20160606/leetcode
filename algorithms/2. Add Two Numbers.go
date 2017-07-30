package algorithms

//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	t := 0
	node := &ListNode{}
	root := node
	var end *ListNode
	for l1 != nil || l2 != nil || t != 0 {
		if l1 != nil {
			t += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t += l2.Val
			l2 = l2.Next
		}
		val := t % 10
		t /= 10
		node.Val = val
		node.Next = &ListNode{}
		end = node
		node = node.Next
	}
	end.Next = nil
	return root
}
