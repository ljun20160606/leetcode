package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	s   string
	t   string
	ans bool
}{

	{"a", "a", true},
	{"aba", "baa", false},
	{"ab", "aa", false},
	{"egg", "add", true},
	{"foo", "bar", false},
	{"paper", "title", true},
}

func Test205(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, isIsomorphic(tc.s, tc.t), "输入:%v", tc)
	}
}
