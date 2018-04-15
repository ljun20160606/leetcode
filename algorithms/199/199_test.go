package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	in  []interface{}
	ans []int
}{

	{
		[]interface{}{},
		nil,
	},

	{
		[]interface{}{1, 2, 3, nil, 5, nil, 4},
		[]int{1, 3, 4},
	},
}

func Test199(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {

		ast.Equal(tc.ans, rightSideView(algorithms.Array2BinaryTree(tc.in)), "输入:%v", tc)
	}
}
