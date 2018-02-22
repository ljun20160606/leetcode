package algorithms

import (
	"math"
)

//Reverse digits of an integer.
//
//Example1: x = 123, return 321
//Example2: x = -123, return -321

func reverse(x int) int {
	flag := 1
	switch {
	case x < -9:
		flag = -1
		x *= -1
	case x < 10:
		return x
	}
	var r int
	for x > 0 {
		r = r*10 + x%10
		x /= 10
	}
	if r > math.MaxInt32 || r < math.MinInt32 {
		return 0
	}
	return flag * r
}
