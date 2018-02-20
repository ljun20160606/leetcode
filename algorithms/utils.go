package algorithms

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func array2BinaryTree(array []interface{}) *TreeNode {
	root := new(*TreeNode)
	insertArrayToTreeNode(root, array, 0)
	return *root
}

func breadthFirstBinaryTree(root *TreeNode, f func(node *TreeNode)) {
	q := []*TreeNode{root}
	for len(q) != 0 {
		n := q[0]
		f(n)
		q = q[1:]
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
}