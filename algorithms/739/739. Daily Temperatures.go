package algorithms

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, len(temperatures))
	top := -1
	for i := range temperatures {
		for top != -1 && temperatures[i] > temperatures[stack[top]] {
			idx := stack[top]
			top--
			res[idx] = i - idx
		}

		top++
		stack[top] = i
	}
	return res
}
