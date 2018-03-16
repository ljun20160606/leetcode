# Single Number

## 描述

```txt
Given an array of integers, every element appears twice except for one. Find that single one.


Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

```

## 题解

```go
package algorithms

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		// n^n == 0
		// a^b^a^b^a == a
		res ^= n
	}
	return res
}

```