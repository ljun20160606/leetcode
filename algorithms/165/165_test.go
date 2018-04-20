package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	version1 string
	version2 string
	ans      int
}{

	{
		"1",
		"1.0.1",
		-1,
	},

	{
		"1.0.1",
		"1",
		1,
	},

	{
		"1",
		"0",
		1,
	},

	{
		"0.1",
		"0.2",
		-1,
	},

	{
		"0.1",
		"1.0",
		-1,
	},

	{
		"1.2",
		"1.1",
		1,
	},

	{
		"1.0",
		"0.1",
		1,
	},

	{
		"1.1",
		"1.1",
		0,
	},
}

func Test165(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, compareVersion(tc.version1, tc.version2), "输入:%v", tc)
	}
}
