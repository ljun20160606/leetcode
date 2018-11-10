package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test344(t *testing.T) {
	ast := assert.New(t)

	qas := []struct {
		in  string
		ans string
	}{
		{
			in:  "hello",
			ans: "olleh",
		},
		{
			in:  "A man, a plan, a canal: Panama",
			ans: "amanaP :lanac a ,nalp a ,nam A",
		},
	}

	for _, v := range qas {
		ast.Equal(v.ans, reverseString(v.in), "输入:%v", v.in)
	}
}
