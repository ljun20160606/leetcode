# Swap Nodes in Pairs

## 描述

```txt

Given a linked list, swap every two adjacent nodes and return its head.



For example,
Given 1-&gt;2-&gt;3-&gt;4, you should return the list as 2-&gt;1-&gt;4-&gt;3.



Your algorithm should use only constant space. You may not modify the values in the list, only nodes itself can be changed.

```

## 题解

```go
package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

//Given a linked list, swap every two adjacent nodes and return its head.
//
//For example,
//Given 1->2->3->4, you should return the list as 2->1->4->3.
//
//Your alsolutionrithm should use only constant space. You may not modify the values in the list, only nodes itself can be changed.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *algorithms.ListNode) *algorithms.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	t := head.Next
	head.Next = swapPairs(t.Next)
	t.Next = head
	return t
}

```