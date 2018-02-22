package algorithms

func BreadthFirstBinaryTree(root *TreeNode, f func(node *TreeNode)) {
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
