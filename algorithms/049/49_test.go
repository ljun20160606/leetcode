package algorithms

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type question struct {
	one []string
	ans [][]string
}

func Test49(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.one
		res := groupAnagrams(p)
		sort.SliceStable(res, func(i, j int) bool {
			return len(res[i]) < len(res[j])
		})
		for i, v := range res {
			sort.Strings(v)
			ast.Equal(a[i], v, "输入:%v", p)
		}
	}
}
