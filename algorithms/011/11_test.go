package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test11(t *testing.T) {
	area := maxArea([]int{4, 9, 8})
	assert.Equal(t, 8, area)
}
