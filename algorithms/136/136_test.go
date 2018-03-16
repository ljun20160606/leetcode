package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 1, 2},
		2,
	},
}

func Test136(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, singleNumber(tc.nums), "输入:%v", tc)
	}
}
