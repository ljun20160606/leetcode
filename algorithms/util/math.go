package util

func MinOfTwo(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func MaxOfTwo(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
