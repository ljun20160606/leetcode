# Minimum Depth of Binary Tree

## 描述

```txt
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
```

## 题解

```go
package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = algorithms.TreeNode

func minDepth(root *TreeNode) int {
	switch {
	case root == nil:
		return 0
	case root.Left == nil:
		return 1 + minDepth(root.Right)
	case root.Right == nil:
		return 1 + minDepth(root.Left)
	default:
		return 1 + algorithms.MinOfTwo(minDepth(root.Left), minDepth(root.Right))
	}
}

```