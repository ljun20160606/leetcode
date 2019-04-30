package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums []int
	k    int
	ans  []int
}{
	{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
}

func Test347(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ans := topKFrequent(tc.nums, tc.k)
		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}