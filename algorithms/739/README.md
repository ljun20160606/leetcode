# Daily Temperatures

## 描述

```txt

Given a list of daily temperatures, produce a list that, for each day in the input, tells you how many days you would have to wait until a warmer temperature.  If there is no future day for which this is possible, put 0 instead.

For example, given the list temperatures = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].


Note:
The length of temperatures will be in the range [1, 30000].
Each temperature will be an integer in the range [30, 100].

```

## 题解

```go
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

		top ++
		stack[top] = i
	}
	return res
}

```