package algorithms

type Interval struct {
	Start int
	End   int
}

type Intervals []Interval

func (in Intervals) Len() int {
	return len(in)
}

func (in Intervals) Less(i, j int) bool {
	x, y := in[i], in[j]
	return x.Start < y.Start
}

func (in Intervals) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}
