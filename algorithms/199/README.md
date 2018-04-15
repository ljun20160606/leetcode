# Binary Tree Right Side View

## 描述

```txt
Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

For example:
Given the following binary tree,

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---


 

You should return [1, 3, 4].

Credits:
Special thanks to @amrsaqr for adding this problem and creating all test cases.

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

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var level []int
	queue := []*TreeNode{root}
	levelLen := len(queue)
	for len(queue) != 0 {
		if levelLen == 0 {
			result = append(result, level[len(level)-1])
			levelLen = len(queue)
			level = []int{}
		}
		t := queue[0]
		queue = queue[1:]
		level = append(level, t.Val)
		levelLen--
		if t.Left != nil {
			queue = append(queue, t.Left)
		}
		if t.Right != nil {
			queue = append(queue, t.Right)
		}
	}
	if len(level) > 0 {
		result = append(result, level[len(level)-1])
	}
	return result
}

```