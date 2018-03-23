# Binary Tree Paths

## 描述

```txt

Given a binary tree, return all root-to-leaf paths.


For example, given the following binary tree:


   1
 /   \
2     3
 \
  5



All root-to-leaf paths are:
[&#34;1-&gt;2-&gt;5&#34;, &#34;1-&gt;3&#34;]


Credits:Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.
```

## 题解

```go
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

```