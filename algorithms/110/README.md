# Balanced Binary Tree

## 描述

```txt
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:


a binary tree in which the depth of the two subtrees of every node never differ by more than 1.


Example 1:

Given the following tree [3,9,20,null,null,15,7]:

    3
   / \
  9  20
    /  \
   15   7

Return true.

Example 2:

Given the following tree [1,2,2,3,3,null,null,4,4]:

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4


Return false.

```

## 题解

```go
package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = algorithms.TreeNode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	diff := heightOfTree(root.Left) - heightOfTree(root.Right)
	return diff >= -1 && diff <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func heightOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return heightOfTree(root.Right) + 1
	}

	if root.Right == nil {
		return heightOfTree(root.Left) + 1
	}

	left := heightOfTree(root.Left)
	right := heightOfTree(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

```