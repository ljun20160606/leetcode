package algorithms

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	one string
	two []string
}

type ans struct {
	one []int
}

func Test30(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{"barfoothefoobarman", []string{"foo", "bar"}},
			ans{[]int{0, 9}},
		},

		{
			para{"barbarfoothefoobarman", []string{"foo", "bar"}},
			ans{[]int{3, 12}},
		},

		{
			para{"attoinattoin", []string{"at", "to", "in"}},
			ans{[]int{0, 2, 4, 6}},
		},

		{
			para{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}},
			ans{[]int{8}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		res := findSubstring(p.one, p.two)
		sort.Ints(res)
		ast.Equal(a.one, res, "输入:%v", p)
	}
}
