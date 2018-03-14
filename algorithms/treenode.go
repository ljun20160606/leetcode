package algorithms

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Array2BinaryTree(array []interface{}) *TreeNode {
	root := new(*TreeNode)
	insertArrayToTreeNode(root, array, 0)
	return *root
}

func insertArrayToTreeNode(n **TreeNode, array []interface{}, index int) {
	if index >= len(array) || array[index] == nil {
		return
	}
	*n = new(TreeNode)
	(*n).Val = array[index].(int)
	insertArrayToTreeNode(&((*n).Left), array, 2*index+1)
	insertArrayToTreeNode(&((*n).Right), array, 2*index+2)
}

// PreIn2Tree 把 preorder 和 inorder 切片转换成 二叉树
func PreIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("preIn2Tree 中两个切片的长度不相等")
	}

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	msg := fmt.Sprintf("%d 不存在于 %v 中", val, nums)
	panic(msg)
}
