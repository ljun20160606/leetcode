package algorithms

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMaxOfTwo(t *testing.T) {
	assert.Equal(t, 1, MaxOfTwo(0, 1))
	assert.Equal(t, 1, MaxOfTwo(1, 0))
}

func TestMinOfTwo(t *testing.T) {
	assert.Equal(t, 1, MinOfTwo(1, 2))
	assert.Equal(t, 1, MinOfTwo(2, 1))
}