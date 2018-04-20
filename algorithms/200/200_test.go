package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	grid [][]byte
	ans  int
}{

	{
		[][]byte{},
		0,
	},

	{
		[][]byte{
			[]byte("111001"),
			[]byte("010001"),
			[]byte("111111"),
		},
		1,
	},

	{
		[][]byte{
			[]byte("11110"),
			[]byte("11010"),
			[]byte("11000"),
			[]byte("00000"),
		},
		1,
	},

	{
		[][]byte{
			[]byte("11000"),
			[]byte("11000"),
			[]byte("00100"),
			[]byte("00011"),
		},
		3,
	},
}

func Test200(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numIslands(tc.grid), "输入:%v", tc)
	}
}
