package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

type TreeNode = algorithms.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	levelLen := len(q)
	var r [][]int
	var level []int
	for len(q) != 0 {
		if levelLen == 0 {
			r = append(r, level)
			levelLen = len(q)
			level = []int{}
		}
		n := q[0]
		levelLen--
		level = append(level, n.Val)
		q = q[1:]
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
	rr := [][]int{level}
	for i := len(r) - 1; i >= 0; i-- {
		rr = append(rr, r[i])
	}
	return rr
}
