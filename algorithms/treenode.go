package algorithms

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
