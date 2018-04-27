package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	ast := assert.New(t)

	stack := new(Stack)
	ast.Equal(true, stack.IsEmpty())
}

func TestStack_Push(t *testing.T) {
	ast := assert.New(t)
	stack := new(Stack)

	ast.NotPanics(func() {
		stack.Push(1)
	})
}

func TestStack_Pop(t *testing.T) {
	ast := assert.New(t)
	stack := new(Stack)

	ast.Panics(func() {
		stack.Pop()
	})

	stack.Push(1)

	ast.Equal(1, stack.Pop())
}
