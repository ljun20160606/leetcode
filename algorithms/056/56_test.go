package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	intervals []algorithms.Interval
}

type ans struct {
	one []algorithms.Interval
}

func Test56(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{},
		{
			para: para{
				intervals: []algorithms.Interval{
					{Start: 2, End: 3},
					{Start: 4, End: 5},
					{Start: 6, End: 7},
					{Start: 8, End: 9},
					{Start: 1, End: 10},
				},
			},
			ans: ans{
				one: []algorithms.Interval{
					{Start: 1, End: 10},
				},
			},
		},
		{
			para: para{intervals: []algorithms.Interval{
				{Start: 1, End: 4},
				{End: 6},
			},
			},
			ans: ans{one: []algorithms.Interval{
				{End: 6},
			},
			},
		},
		{
			para: para{intervals: []algorithms.Interval{
				{Start: 1, End: 4},
				{Start: 2, End: 5},
			},
			},
			ans: ans{one: []algorithms.Interval{
				{Start: 1, End: 5},
			},
			},
		},
		{
			para: para{intervals: []algorithms.Interval{
				{Start: 1, End: 3},
				{Start: 2, End: 6},
				{Start: 8, End: 10},
				{Start: 15, End: 18},
			}},
			ans: ans{one: []algorithms.Interval{
				{Start: 1, End: 6},
				{Start: 8, End: 10},
				{Start: 15, End: 18},
			},
			},
		},
		{
			para: para{intervals: []algorithms.Interval{
				{Start: 1, End: 6},
			},
			},
			ans: ans{one: []algorithms.Interval{
				{Start: 1, End: 6},
			},
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, merge(p.intervals), "输入:%v", p)
	}
}
