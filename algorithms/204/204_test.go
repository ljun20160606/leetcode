package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	n   int
	ans int
}{

	{
		499979,
		41537,
	},

	{
		2,
		0,
	},

	{
		5,
		2,
	},

	{
		10,
		4,
	},

	{
		100,
		25,
	},
}

func Test204(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, countPrimes(tc.n), "输入:%v", tc)
	}
}
