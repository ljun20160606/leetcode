package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

//Reverse a singly linked list.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *algorithms.ListNode) *algorithms.ListNode {
	var newHead, tempHead *algorithms.ListNode
	for head != nil {
		tempHead = head.Next
		head.Next = newHead
		newHead = head
		head = tempHead
	}
	return newHead
}
