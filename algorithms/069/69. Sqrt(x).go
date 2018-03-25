package algorithms

func mySqrt(x int) int {
	res := x

	// 牛顿迭代法
	// https://baike.baidu.com/item/%E7%89%9B%E9%A1%BF%E8%BF%AD%E4%BB%A3%E6%B3%95
	// https://en.wikipedia.org/wiki/Newton%27s_method
	for res*res > x {
		res = (res + x/res) / 2
	}

	return res
}
