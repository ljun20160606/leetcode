# Same Tree

## 描述

```txt

Given two binary trees, write a function to check if they are the same or not.


Two binary trees are considered the same if they are structurally identical and the nodes have the same value.




Example 1:
Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true



Example 2:
Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false



Example 3:
Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false


```

## 题解

```go
package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

//Given two binary trees, write a function to check if they are the same or not.
//
//Two binary trees are considered the same if they are structurally identical and the nodes have the same value.
//
//
//	Example 1:
//
//	Input:   1         1
//          / \       / \
//		   2   3     2   3
//
//	[1,2,3],   [1,2,3]
//
//	Output: true
//	Example 2:
//
//	Input:   1         1
//          /           \
//		   2             2
//
//  [1,2],     [1,null,2]
//
//  Output: false
//  Example 3:
//
//	Input:   1         1
///           \       / \
//		   2   1     1   2
//
//  [1,2,1],   [1,1,2]
//
//  Output: false

func isSameTree(p *algorithms.TreeNode, q *algorithms.TreeNode) bool {
	if p == q {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

```