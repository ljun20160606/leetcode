# Sqrt(x)

## 描述

```txt
Implement int sqrt(int x).

Compute and return the square root of x.

x is guaranteed to be a non-negative integer.



Example 1:
Input: 4
Output: 2



Example 2:
Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we want to return an integer, the decimal part will be truncated.


```

## 题解

```go
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

```