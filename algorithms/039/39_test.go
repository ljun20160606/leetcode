package algrithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	parameter
	answer
}

type parameter struct {
	candidates []int
	target     int
}

type answer struct {
	one [][]int
}

func Test39(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			parameter{[]int{7, 3, 2}, 18},
			answer{[][]int{
				{2, 2, 2, 2, 2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2, 2, 3, 3},
				{2, 2, 2, 2, 3, 7},
				{2, 2, 2, 3, 3, 3, 3},
				{2, 2, 7, 7},
				{2, 3, 3, 3, 7},
				{3, 3, 3, 3, 3, 3},
			}},
		},
		{
			parameter{[]int{2, 3, 6, 7}, 7},
			answer{[][]int{
				{2, 2, 3},
				{7},
			}},
		},
		{
			parameter{[]int{8, 7, 4, 3}, 11},
			answer{[][]int{
				{3, 4, 4},
				{3, 8},
				{4, 7},
			}},
		},
	}

	for _, q := range qs {
		a, p := q.answer, q.parameter

		ast.Equal(a.one, combinationSum(p.candidates, p.target), "输入:%v", p)
	}
}
