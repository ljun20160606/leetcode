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
	nums []int
	k    int
}

type ans struct {
	one bool
}

func Test219(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{
				[]int{1, 0, 1, 1},
				1,
			},
			ans{
				true,
			},
		},

		{
			para{
				[]int{1, 2, 3, 4, 5},
				0,
			},
			ans{
				false,
			},
		},

		{
			para{
				[]int{1, 2, 3, 4, 5},
				2,
			},
			ans{
				false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, containsNearbyDuplicate(p.nums, p.k), "输入:%v", p)
	}
}
