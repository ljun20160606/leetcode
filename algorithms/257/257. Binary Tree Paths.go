package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"strconv"
	"strings"
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

func binaryTreePaths(root *TreeNode) []string {
	var result []string
	helper(root, &result, []string{})
	return result
}

func helper(root *TreeNode, result *[]string, temp []string) {
	if root == nil {
		return
	}
	temp = append(temp, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		*result = append(*result, strings.Join(temp, "->"))
		return
	}
	helper(root.Left, result, temp)
	helper(root.Right, result, temp)
}
