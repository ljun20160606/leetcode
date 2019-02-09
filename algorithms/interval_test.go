package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartIntervals_Len(t *testing.T) {
	var intervals IntervalsStart
	assert.Equal(t, 0, intervals.Len())
}

func TestStartIntervals_Less(t *testing.T) {
	intervals := IntervalsStart{{0, 1}, {1, 2}}
	assert.True(t, intervals.Less(0, 1))
}

func TestStartIntervals_Swap(t *testing.T) {
	intervals := IntervalsStart{{0, 1}, {1, 2}}
	assert.NotPanics(t, func() {
		intervals.Swap(0, 1)
		assert.Equal(t, intervals[0], Interval{1, 2})
		assert.Equal(t, intervals[1], Interval{0, 1})
	})
}

func TestEndIntervals_Len(t *testing.T) {
	var intervals IntervalsEnd
	assert.Equal(t, 0, intervals.Len())
}

func TestEndIntervals_Less(t *testing.T) {
	intervals := IntervalsEnd{{0, 1}, {1, 2}}
	assert.True(t, intervals.Less(0, 1))
}

func TestEndIntervals_Swap(t *testing.T) {
	intervals := IntervalsEnd{{0, 1}, {1, 2}}
	assert.NotPanics(t, func() {
		intervals.Swap(0, 1)
		assert.Equal(t, intervals[0], Interval{1, 2})
		assert.Equal(t, intervals[1], Interval{0, 1})
	})
}

func TestInt2sToIntervals(t *testing.T) {
	int2s := [][]int{{2, 3}}
	assert.Equal(t, []Interval{{2, 3}}, Int2sToIntervals(int2s))
}
