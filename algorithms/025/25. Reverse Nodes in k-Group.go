package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

//Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
//
//k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
//
//You may not alter the values in the nodes, only nodes itself may be changed.
//
//Only constant memory is allowed.
//
//For example,
//Given this linked list: 1->2->3->4->5
//
//For k = 2, you should return: 2->1->4->3->5
//
//For k = 3, you should return: 3->2->1->4->5

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *algorithms.ListNode, k int) *algorithms.ListNode {
	start := head
	for i := 0; i < k; i++ {
		if head == nil {
			return start
		}
		head = head.Next
	}
	newHead := reverseGroup(start, head)
	start.Next = reverseKGroup(head, k)
	return newHead

}

func reverseGroup(start, end *algorithms.ListNode) *algorithms.ListNode {
	var newHead, tempHead *algorithms.ListNode
	for start != end {
		tempHead = start.Next
		start.Next = newHead
		newHead = start
		start = tempHead
	}
	return newHead
}
