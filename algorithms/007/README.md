# Reverse Integer

## 描述

```txt
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
Input: 123
Output:  321



Example 2:
Input: -123
Output: -321



Example 3:
Input: 120
Output: 21



Note:
Assume we are dealing with an environment which could only hold integers within the 32-bit signed integer range. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

```

## 题解

```go
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

```