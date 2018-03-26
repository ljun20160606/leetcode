package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test125(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		s   string
		ans bool
	}{

		{
			"0p",
			false,
		},

		{
			"0",
			true,
		},

		{
			"race a car",
			false,
		},

		{
			"A man, a plan, a canal: Panama",
			true,
		},
	}

	for _, tc := range tcs {

		ast.Equal(tc.ans, isPalindrome(tc.s), "输入:%v", tc)
	}
}
