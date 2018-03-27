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

func levelOrder(root *TreeNode) [][]int {
	var r [][]int
	helper(&r, root, 0)
	return r
}

func helper(r *[][]int, node *TreeNode, level int) {
	if node == nil {
		return
	}

	if level >= len(*r) {
		*r = append(*r, []int{})
	}

	(*r)[level] = append((*r)[level], node.Val)

	i := level + 1
	helper(r, node.Left, i)
	helper(r, node.Right, i)
}
