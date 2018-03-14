# Binary Tree Inorder Traversal

## 描述

```txt
Given a binary tree, return the inorder traversal of its nodes&#39; values.


For example:
Given binary tree [1,null,2,3],
   1
    \
     2
    /
   3



return [1,3,2].


Note: Recursive solution is trivial, could you do it iteratively?
```

## 题解

```go
package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

//Given a binary tree, return the inorder traversal of its nodes' values.
//
//For example:
//Given binary tree [1,null,2,3],
//1
//\
//2
///
//3
//return [1,3,2].
//
//Note: Recursive solution is trivial, could you do it iteratively?

func inorderTraversal(root *algorithms.TreeNode) (r []int) {
	inorderTraversalHelper(root, &r)
	return
}

func inorderTraversalHelper(root *algorithms.TreeNode, r *[]int) {
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

```