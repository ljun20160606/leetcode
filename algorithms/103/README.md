# Binary Tree Zigzag Level Order Traversal

## 描述

```txt
Given a binary tree, return the zigzag level order traversal of its nodes&#39; values. (ie, from left to right, then right to left for the next level and alternate between).


For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7



return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]


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

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		// 出现了新的 level
		if level >= len(res) {
			res = append(res, []int{})
		}
		if level%2 == 0 {
			res[level] = append(res[level], root.Val)
		} else {
			temp := make([]int, len(res[level])+1)
			temp[0] = root.Val
			copy(temp[1:], res[level])
			res[level] = temp
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}

```