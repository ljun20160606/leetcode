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
	one int
}

type ans struct {
	one string
}

func Test38(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{},

		{
			para{1},
			ans{"1"},
		},

		{
			para{2},
			ans{"11"},
		},

		{
			para{3},
			ans{"21"},
		},

		{
			para{4},
			ans{"1211"},
		},

		{
			para{5},
			ans{"111221"},
		},

		{
			para{6},
			ans{"312211"},
		},

		{
			para{7},
			ans{"13112221"},
		},

		{
			para{20},
			ans{"11131221131211132221232112111312111213111213211231132132211211131221131211221321123113213221123113112221131112311332211211131221131211132211121312211231131112311211232221121321132132211331121321231231121113112221121321133112132112312321123113112221121113122113121113123112112322111213211322211312113211"},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, countAndSay(p.one), "输入:%v", p)
	}
}
