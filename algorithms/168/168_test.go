package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	n   int
	ans string
}{

	{
		100,
		"CV",
	},

	{
		1000,
		"ALL",
	},

	{
		53,
		"BA",
	},

	{
		52,
		"AZ",
	},

	{
		26,
		"Z",
	},

	{
		1,
		"A",
	},

	{
		28,
		"AB",
	},
}

func Test168(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, convertToTitle(tc.n), "输入:%v", tc)
	}
}
