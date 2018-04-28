package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	temperatures []int
	ans          []int
}{
	{
		[]int{71, 71, 71},
		[]int{0, 0, 0},
	},

	{
		[]int{73, 74, 75, 71, 69, 72, 76, 73},
		[]int{1, 1, 4, 2, 1, 1, 0, 0},
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, dailyTemperatures(tc.temperatures), "输入:%v", tc)
	}
}
