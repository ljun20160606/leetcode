package other

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKnapsack(t *testing.T) {
	assert.Equal(t, 26, knapsack([]item{
		{
			volume: 2,
			price:  6,
		},
		{
			volume: 5,
			price:  5,
		},
		{
			volume: 3,
			price:  4,
		},
		{
			volume: 10,
			price:  15,
		},
		{
			volume: 2,
			price:  5,
		},
	}, 15))
}
