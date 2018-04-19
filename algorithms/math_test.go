package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxOfTwo(t *testing.T) {
	assert.Equal(t, 1, MaxOfTwo(0, 1))
	assert.Equal(t, 1, MaxOfTwo(1, 0))
}

func TestMinOfTwo(t *testing.T) {
	assert.Equal(t, 1, MinOfTwo(1, 2))
	assert.Equal(t, 1, MinOfTwo(2, 1))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 1, Abs(-1))
	assert.Equal(t, 1, Abs(1))
}
