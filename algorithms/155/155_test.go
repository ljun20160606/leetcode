package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test155(t *testing.T) {
	ast := assert.New(t)

	s := Constructor()

	s.Push(-2)
	// [-2]
	s.Push(0)
	// [-2, 0]
	s.Push(-3)
	// [-2, 0, -3]
	ast.Equal(-3, s.GetMin(), "get min from [-2, 0, -3]")
	// [-2, 0, -3]
	s.Pop()
	// [-2, 0]
	ast.Equal(0, s.Top(), "get top from [-2, 0]")
	// [-2, 0]
	ast.Equal(-2, s.GetMin(), "get min from [-2, 0]")
	// [-2, 0]
	s.Pop()
	// [-2]
	ast.Equal(-2, s.Top())
	// [-2]
	ast.Equal(-2, s.GetMin())
	// [-2]
	s.Pop()
}
