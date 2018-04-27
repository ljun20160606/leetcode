package algorithms

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ljun20160606/leetcode/algorithms"
)

var tcs = []struct {
	in     []interface{}
	expect []int
}{
	{
		[]interface{}{1, 2, 3},
		[]int{2, 3, 1},
	},

	{
		[]interface{}{1, nil, 2, nil, nil, 3},
		[]int{3, 2, 1},
	},
}

func Test145(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {

		ast.Equal(tc.expect, postorderTraversal(algorithms.Array2BinaryTree(tc.in)), "输入:%v", tc)
		ast.Equal(tc.expect, iteration(algorithms.Array2BinaryTree(tc.in)), "输入:%v", tc)
	}
}
