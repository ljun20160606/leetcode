# Remove Linked List Elements

## 描述

```txt
Remove all elements from a linked list of integers that have value val.

Example
Given: 1 --> 2 --> 6 --> 3 --> 4 --> 5 --> 6,  val = 6
Return: 1 --> 2 --> 3 --> 4 --> 5


Credits:Special thanks to @mithmatt for adding this problem and creating all test cases.
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

```