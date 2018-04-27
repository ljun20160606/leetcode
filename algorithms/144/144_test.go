package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	in     []interface{}
	expect []int
}{
	{
		[]interface{}{1, 2, 3},
		[]int{1, 2, 3},
	},

	{
		[]interface{}{1, nil, 2, nil, nil, 3},
		[]int{1, 2, 3},
	},
}

func Test144(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {

		ast.Equal(tc.expect, preorderTraversal(algorithms.Array2BinaryTree(tc.in)), "输入:%v", tc)
		ast.Equal(tc.expect, iteration(algorithms.Array2BinaryTree(tc.in)), "输入:%v", tc)
	}
}
