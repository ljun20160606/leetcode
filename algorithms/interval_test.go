package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntervals_Len(t *testing.T) {
	var intervals IntervalsStart
	assert.Equal(t, 0, intervals.Len())
}

func TestIntervals_Less(t *testing.T) {
	intervals := IntervalsStart{{0, 1}, {1, 2}}
	assert.True(t, intervals.Less(0, 1))
}

func TestIntervals_Swap(t *testing.T) {
	intervals := IntervalsStart{{0, 1}, {1, 2}}
	assert.NotPanics(t, func() {
		intervals.Swap(0, 1)
		assert.Equal(t, intervals[0], Interval{1, 2})
		assert.Equal(t, intervals[1], Interval{0, 1})
	})
}
