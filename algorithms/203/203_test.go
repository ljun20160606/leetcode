package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	ints []int
	val  int
	ans  []int
}{

	{
		[]int{1, 2, 6, 3, 4, 5, 6},
		6,
		[]int{1, 2, 3, 4, 5},
	},
}

func Test_removeElements(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(algorithms.Array2ListNode(tc.ans), removeElements(algorithms.Array2ListNode(tc.ints), tc.val), "输入:%v", tc)
	}
}
