package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one string
	two int
}

type ans struct {
	one string
}

type question struct {
	p para
	a ans
}

func Test6(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
				two: 1,
			},
			a: ans{
				one: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			},
		},
		{
			p: para{
				one: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
				two: 2,
			},
			a: ans{
				one: "ACEGIKMOQSUWYBDFHJLNPRTVXZ",
			},
		},
		{
			p: para{
				one: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
				two: 3,
			},
			a: ans{
				one: "AEIMQUYBDFHJLNPRTVXZCGKOSW",
			},
		},
		{
			p: para{
				one: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
				two: 4,
			},
			a: ans{
				one: "AGMSYBFHLNRTXZCEIKOQUWDJPV",
			},
		},
		{
			p: para{
				one: "ABCDEFGHIJKLMNOPQRSTUVWX",
				two: 5,
			},
			a: ans{
				one: "AIQBHJPRXCGKOSWDFLNTVEMU",
			},
		},
		{
			p: para{
				one: "A",
				two: 3,
			},
			a: ans{
				one: "A",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, convert(p.one, p.two), "输入:%v", p)
	}
}
