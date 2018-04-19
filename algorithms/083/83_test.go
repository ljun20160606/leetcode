package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test83(t *testing.T) {
	ast := assert.New(t)

	qst := []struct {
		in       []int
		expected []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1, 1, 2},
			[]int{1, 2},
		},
		{
			[]int{1, 1, 2, 3, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 1, 1},
			[]int{1},
		},
	}

	for i := range qst {
		ast.Equal(algorithms.Array2ListNode(qst[i].expected), deleteDuplicates(algorithms.Array2ListNode(qst[i].in)), "输入 %v", qst[i].in)
	}
}
