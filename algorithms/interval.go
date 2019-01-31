package algorithms

type Interval struct {
	Start int
	End   int
}

type IntervalsStart []Interval

func (in IntervalsStart) Len() int {
	return len(in)
}

func (in IntervalsStart) Less(i, j int) bool {
	x, y := in[i], in[j]
	return x.Start < y.Start
}

func (in IntervalsStart) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

type IntervalsEnd []Interval

func (in IntervalsEnd) Len() int {
	return len(in)
}

func (in IntervalsEnd) Less(i, j int) bool {
	x, y := in[i], in[j]
	return x.End < y.End
}

func (in IntervalsEnd) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func Int2sToIntervals(v [][]int) []Interval {
	var intervals []Interval
	for i := range v {
		ints := v[i]
		intervals = append(intervals, Interval{
			Start: ints[0],
			End:   ints[1],
		})
	}

	return intervals
}
