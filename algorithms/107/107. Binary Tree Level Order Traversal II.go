package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

type TreeNode = algorithms.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var level []int
	queue := []*TreeNode{root}
	levelLen := len(queue)
	for len(queue) != 0 {
		if levelLen == 0 {
			result = append(result, level)
			levelLen = len(queue)
			level = []int{}
		}
		n := queue[0]
		queue = queue[1:]
		level = append(level, n.Val)
		levelLen--
		if n.Left != nil {
			queue = append(queue, n.Left)
		}
		if n.Right != nil {
			queue = append(queue, n.Right)
		}
	}
	rr := [][]int{level}
	for i := len(result) - 1; i >= 0; i-- {
		rr = append(rr, result[i])
	}
	return rr
}
