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
	path string
}

type ans struct {
	one string
}

func Test71(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{
				"/home/",
			},
			ans{
				"/home",
			},
		},

		{
			para{
				"/a/./b/../../c/",
			},
			ans{
				"/c",
			},
		},

		{
			para{
				"/../",
			},
			ans{
				"/",
			},
		},

		{
			para{
				"/home//foo/",
			},
			ans{
				"/home/foo",
			},
		},

		{
			para{
				"/home/./foo/",
			},
			ans{
				"/home/foo",
			},
		},

		{
			para{
				"/home////////////////////////foo/",
			},
			ans{
				"/home/foo",
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, simplifyPath(p.path), "输入:%v", p)
	}
}
