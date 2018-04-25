# Rotate List

## 描述

```txt
Given a linked list, rotate the list to the right by k places, where k is non-negative.

Example 1:

Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL


Example 2:

Input: 0->1->2->NULL, k = 4
Output: 2->0->1->NULL
Explanation:
rotate 1 steps to the right: 2->0->1->NULL
rotate 2 steps to the right: 1->2->0->NULL
rotate 3 steps to the right: 0->1->2->NULL
rotate 4 steps to the right: 2->0->1->NULL

```

## 题解

```go
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

```