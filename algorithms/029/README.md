# Divide Two Integers

## 描述

```txt

Divide two integers without using multiplication, division and mod operator.


If it is overflow, return MAX_INT.

```

## 题解

```go
package algorithms

import "math"

//Divide two integers without using multiplication, division and mod operator.
//
//If it is overflow, return MAX_INT.

// dividend = divisor * 2^n1 + divisor * 2^n2 + ... + divisor * 2^nk
// divisor 一直乘2直到比 分子大求出 n1
// dividend 减去2^n1 作为下次的 divisor
// 同理求出 n2 n3 nk
// 结果为n1+n2+n3+...+nk
func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return math.MaxInt32
	}
	flag := 1
	if dividend < 0 {
		flag *= -1
		dividend *= -1
	}
	if divisor < 0 {
		flag *= -1
		divisor *= -1
	}
	var multiple int
	for divisor <= dividend {
		t := divisor // t = divisor
		count := 1
		for t<<1 <= dividend {
			t <<= 1     // divisor * 2
			count <<= 1 // 倍数 * 2
		}
		multiple += count
		dividend -= t
	}
	if multiple >= math.MaxInt32 {
		if flag == 1 {
			return math.MaxInt32
		}
		return math.MinInt32
	}
	return multiple * flag
}

```