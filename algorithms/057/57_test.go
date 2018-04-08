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
	intervals   []Interval
	newInterval Interval
}

type ans struct {
	one []Interval
}

func Test57(t *testing.T) {
	ast := assert.New(t)
	qs := []question{

		{
			para{[]Interval{
				{1, 3},
				{6, 9},
			},
				Interval{7, 8},
			},
			ans{[]Interval{
				{1, 3},
				{6, 9},
			},
			},
		},

		{
			para{[]Interval{
				{1, 3},
				{6, 9},
			},
				Interval{4, 5},
			},
			ans{[]Interval{
				{1, 3},
				{4, 5},
				{6, 9},
			},
			},
		},

		{
			para{[]Interval{
				{1, 3},
				{6, 9},
			},
				Interval{2, 5},
			},
			ans{[]Interval{
				{1, 5},
				{6, 9},
			},
			},
		},

		{
			para{[]Interval{
				{1, 3},
				{6, 9},
			},
				Interval{-1, 0},
			},
			ans{[]Interval{
				{-1, 0},
				{1, 3},
				{6, 9},
			},
			},
		},

		{
			para{[]Interval{
				{1, 3},
				{6, 9},
			},
				Interval{11, 20},
			},
			ans{[]Interval{
				{1, 3},
				{6, 9},
				{11, 20},
			},
			},
		},

		{
			para{[]Interval{},
				Interval{2, 5},
			},
			ans{[]Interval{
				{2, 5},
			},
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, insert(p.intervals, p.newInterval), "输入:%v", p)
	}
}
