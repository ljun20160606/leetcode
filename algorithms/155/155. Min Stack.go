package algorithms

type MinStack struct {
	stack []item
}

type item struct {
	min, x int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(x int) {
	min := x
	if len(m.stack) > 0 && m.GetMin() < x {
		min = m.GetMin()
	}
	m.stack = append(m.stack, item{min: min, x: x})
	return
}

func (m *MinStack) Pop() {
	length := len(m.stack)
	m.stack = m.stack[:length-1]
	return
}

func (m *MinStack) Top() int {
	length := len(m.stack)
	return m.stack[length-1].x
}

func (m *MinStack) GetMin() int {
	length := len(m.stack)
	return m.stack[length-1].min
}

//type MinStack struct {
//	stack []int
//	min   []int
//}
//
///** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{}
//}
//
//func (m *MinStack) Push(x int) {
//	m.stack = append(m.stack, x)
//	if len(m.min) == 0 {
//		m.min = append(m.min, len(m.stack)-1)
//		return
//	}
//	min := m.stack[m.min[len(m.min)-1]]
//	if x <= min {
//		m.min = append(m.min, len(m.stack)-1)
//	}
//}
//
//func (m *MinStack) Pop() {
//	if len(m.stack) == 0 {
//		return
//	}
//	i := len(m.stack) - 1
//	m.stack = m.stack[:i]
//	if len(m.min) == 0 {
//		return
//	}
//	i2 := len(m.min) - 1
//	if m.min[i2] == i {
//		m.min = m.min[:i2]
//	}
//}
//
//func (m *MinStack) Top() int {
//	return m.stack[len(m.stack)-1]
//}
//
//func (m *MinStack) GetMin() int {
//	if len(m.min) > 0 {
//		return m.stack[m.min[len(m.min)-1]]
//	}
//	return m.Top()
//}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
