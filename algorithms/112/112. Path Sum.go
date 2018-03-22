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

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if sum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}
	val := sum - root.Val
	return hasPathSum(root.Left, val) || hasPathSum(root.Right, val)
}
