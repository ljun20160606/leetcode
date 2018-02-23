package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
}

type ans struct {
	one []int
}

func Test31(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{[]int{5, 1, 1}},
			ans{[]int{1, 1, 5}},
		},

		{
			para{[]int{1, 1}},
			ans{[]int{1, 1}},
		},

		{
			para{[]int{1, 5, 1}},
			ans{[]int{5, 1, 1}},
		},

		{
			para{[]int{1, 5, 4, 3, 2}},
			ans{[]int{2, 1, 3, 4, 5}},
		},

		{
			para{[]int{1, 2, 7, 4, 3, 1}},
			ans{[]int{1, 3, 1, 2, 4, 7}},
		},

		{
			para{[]int{1, 2, 3}},
			ans{[]int{1, 3, 2}},
		},

		{
			para{[]int{3, 2, 1}},
			ans{[]int{1, 2, 3}},
		},

		{
			para{[]int{1, 1, 5}},
			ans{[]int{1, 5, 1}},
		},

		{
			para{[]int{2, 1}},
			ans{[]int{1, 2}},
		},

		{
			para{[]int{1}},
			ans{[]int{1}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		copyOne := make([]int, len(p.one))
		copy(copyOne, p.one)
		nextPermutation(p.one)
		ast.Equal(a.one, p.one, "输入:%v", p)

		nextPermutationBest(copyOne)
		ast.Equal(a.one, copyOne, "输入:%v", p)
	}
}
