package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test12(t *testing.T) {
	roman := intToRoman(112)
	assert.Equal(t, "CXII", roman)
}
