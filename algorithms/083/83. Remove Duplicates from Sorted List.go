package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = algorithms.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tmp := head
	for {
		if tmp.Next == nil {
			break
		}
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return head
}
