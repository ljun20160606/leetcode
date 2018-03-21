package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test110(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		in  []interface{}
		ans bool
	}{

		{
			[]interface{}{3, 9, 20, nil, nil, 15, 7},
			true,
		},

		{
			[]interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4},
			false,
		},

		{
			[]interface{}{1, 1, 1, nil, nil, nil, 1},
			true,
		},

		{
			[]interface{}{1, 1, 1, 1, nil},
			true,
		},
	}

	for _, tc := range tcs {

		ast.Equal(tc.ans, isBalanced(algorithms.Array2BinaryTree(tc.in)), "输入:%v", tc)
	}
}
