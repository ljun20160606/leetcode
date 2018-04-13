package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

// para 是参数
type para struct {
	nums []int
}

// ans 是答案
type ans struct {
	one bool
}

func Test217(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{
				[]int{1, 2, 3, 3},
			},
			ans{
				true,
			},
		},
		{
			para{
				[]int{1, 2, 3, 4},
			},
			ans{
				false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, containsDuplicate(p.nums), "输入:%v", p)
	}
}
