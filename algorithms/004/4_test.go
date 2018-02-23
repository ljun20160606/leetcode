package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindKth(t *testing.T) {
	assert.Equal(t, float64(3), findKth([]int{1, 2, 3, 4}, []int{2, 3, 4, 7, 9}, 3))
}

type parameter struct {
	one []int
	two []int
}

type answer struct {
	one float64
}

type question struct {
	parameter
	answer
}

func Test4(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			parameter: parameter{
				one: []int{1, 3},
				two: []int{2},
			},
			answer: answer{
				one: 2,
			},
		},
		{
			parameter: parameter{
				one: []int{1, 3},
				two: []int{2, 4},
			},
			answer: answer{
				one: 2.5,
			},
		},
	}

	for _, q := range qs {
		a, p := q.answer, q.parameter
		ast.Equal(a.one, findMedianSortedArrays(p.one, p.two), "输入:%v", p)
	}

	ast.Panics(func() { findMedianSortedArrays([]int{}, []int{}) }, "对空切片求中位数，却没有panic")
}
