package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8

func addTwoNumbers(l1 *algorithms.ListNode, l2 *algorithms.ListNode) *algorithms.ListNode {
	t := 0
	head := new(*algorithms.ListNode)
	current := head
	for l1 != nil || l2 != nil || t != 0 {
		if l1 != nil {
			t += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t += l2.Val
			l2 = l2.Next
		}
		if *current == nil {
			*current = new(algorithms.ListNode)
		}
		val := t % 10
		t /= 10
		(**current).Val = val
		current = &((*current).Next)
	}
	return *head
}
