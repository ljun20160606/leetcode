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

func removeElements(head *ListNode, val int) *ListNode {
	headPre := ListNode{Next: head}

	temp := &headPre
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	return headPre.Next
}
