package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test104(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		pre []interface{}
		ans int
	}{
		{
			[]interface{}{},
			0,
		},

		{
			[]interface{}{3, 9, 20, 15, 7},
			3,
		},

		{
			[]interface{}{3, 9, 20, 15, 10, 7},
			3,
		},
	}

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxDepth(algorithms.Array2BinaryTree(tc.pre)), "输入:%v", tc)
	}
}
