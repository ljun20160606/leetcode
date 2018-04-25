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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	var list []*ListNode
	temp := head
	for {
		list = append(list, temp)
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}

	rotateCount := k % len(list)
	if rotateCount == 0 {
		return head
	}
	pivotIndex := len(list) - rotateCount - 1
	pivot := list[pivotIndex]
	newHead := pivot.Next
	pivot.Next = nil
	temp.Next = head
	return newHead
}
