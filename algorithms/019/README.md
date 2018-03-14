# Remove Nth Node From End of List

## 描述

```txt
Given a linked list, remove the nth node from the end of list and return its head.

For example,

   Given linked list: 1-&gt;2-&gt;3-&gt;4-&gt;5, and n = 2.

   After removing the second node from the end, the linked list becomes 1-&gt;2-&gt;3-&gt;5.


Note:
Given n will always be valid.
Try to do this in one pass.

```

## 题解

```go
package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

//Given a linked list, remove the nth node from the end of list and return its head.
//
//For example,
//
//Given linked list: 1->2->3->4->5, and n = 2.
//
//After removing the second node from the end, the linked list becomes 1->2->3->5.

//Note:
//Given n will always be valid.
//Try to do this in one pass.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 双指针 first 先走n个节点, second 再走，保证了他们永远间隔n-1, 当 first 走到头时, second代表倒数第n个
// 边界条件 当 second 为第一个节点时，返回第二个节点即可
func removeNthFromEnd(head *algorithms.ListNode, n int) *algorithms.ListNode {
	first, second := head, head
	for i := n; i > 0; i-- {
		first = first.Next
	}
	var prev *algorithms.ListNode
	for first != nil {
		first = first.Next
		prev = second
		second = second.Next
	}
	if prev == nil {
		t := head.Next
		head.Next = nil
		head = t
	} else {
		prev.Next = second.Next
		second.Next = nil
	}
	return head
}

```