# Binary Tree Postorder Traversal

## 描述

```txt
Given a binary tree, return the postorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [3,2,1]


Follow up: Recursive solution is trivial, could you do it iteratively?

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

func postorderTraversal(root *TreeNode) []int {
	var result []int
	helper(&result, root)
	return result
}

func helper(result *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	helper(result, node.Left)
	helper(result, node.Right)
	*result = append(*result, node.Val)
}

func iteration(root *TreeNode) []int {
	var reverse []int
	stack := new(algorithms.Stack)
	stack.Push(root)
	for !stack.IsEmpty() {
		node := stack.Pop().(*TreeNode)

		reverse = append(reverse, node.Val)

		if node.Left != nil {
			stack.Push(node.Left)
		}

		if node.Right != nil {
			stack.Push(node.Right)
		}
	}
	result := make([]int, len(reverse))
	for i := range reverse {
		result[len(reverse)-1-i] = reverse[i]
	}
	return result
}

```