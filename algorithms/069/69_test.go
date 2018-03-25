package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test69(t *testing.T) {
	ast := assert.New(t)

	qst := []struct {
		in  int
		out int
	}{
		{
			8,
			2,
		},
	}

	for _, q := range qst {
		ast.Equal(q.out, mySqrt(q.in), "输入:%v", q.in)
	}
}
