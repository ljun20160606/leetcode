package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

type TreeNode = algorithms.TreeNode

func inorderTraversal(root *TreeNode) (r []int) {
	inorderTraversalHelper(root, &r)
	return
}

func inorderTraversalHelper(root *TreeNode, r *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorderTraversalHelper(root.Left, r)
	}
	*r = append(*r, root.Val)
	if root.Right != nil {
		inorderTraversalHelper(root.Right, r)
	}
}
