package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Problem0108(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		nums []int
		tree []interface{}
	}{

		{
			[]int{},
			nil,
		},

		{
			[]int{-10, -3, 0, 5, 9},
			[]interface{}{0, -3, 9, -10, nil, 5},
		},
	}

	for _, tc := range tcs {

		ast.Equal(algorithms.Array2BinaryTree(tc.tree), sortedArrayToBST(tc.nums), "输入:%v", tc)
	}
}
