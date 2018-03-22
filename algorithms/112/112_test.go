package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test112(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		in  []interface{}
		sum int
		ans bool
	}{

		{
			[]interface{}{1, 2},
			1,
			false,
		},

		{
			[]interface{}{},
			0,
			false,
		},

		{
			[]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1},
			22,
			true,
		},
	}

	for _, tc := range tcs {
		ast.Equal(tc.ans, hasPathSum(algorithms.Array2BinaryTree(tc.in), tc.sum), "输入:%v", tc)
	}
}
