# Symmetric Tree

## 描述

```txt
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).


For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
    1
   / \
  2   2
 / \ / \
3  4 4  3



But the following [1,2,2,null,3,null,3]  is not:
    1
   / \
  2   2
   \   \
   3    3




Note:
Bonus points if you could solve it both recursively and iteratively.

```

## 题解

```go
package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

//Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
//
//For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
//
//1
/// \
//2   2
/// \ / \
//3  4 4  3
//But the following [1,2,2,null,3,null,3] is not:
//1
/// \
//2   2
//\   \
//3    3
//Note:
//Bonus points if you could solve it both recursively and iteratively.

func isSymmetric(root *algorithms.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(p *algorithms.TreeNode, q *algorithms.TreeNode) bool {
	if p == q {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSymmetricTree(p.Left, q.Right) && isSymmetricTree(p.Right, q.Left)
}

```