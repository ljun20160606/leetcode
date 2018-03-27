package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test12(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		in  []interface{}
		ans [][]int
	}{
		{
			[]interface{}{3, 9, 20, nil, nil, 15, 7},
			[][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
	}

	for _, tc := range tcs {
		ast.Equal(tc.ans, levelOrder(algorithms.Array2BinaryTree(tc.in)), "输入:%v", tc)
	}
}
