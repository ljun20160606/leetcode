package algorithms

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(i interface{}) {
	*s = append(*s, i)
}

func (s *Stack) Pop() interface{} {
	end := len(*s) - 1
	r := (*s)[end]
	*s = (*s)[:end]
	return r
}

func (s *Stack) Peek() interface{} {
	return (*s)[len(*s)-1]
}
