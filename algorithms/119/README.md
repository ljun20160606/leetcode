# Pascal's Triangle II

## 描述

```txt
Given an index k, return the kth row of the Pascal's triangle.


For example, given k = 3,
Return [1,3,3,1].



Note:
Could you optimize your algorithm to use only O(k) extra space?

```

## 题解

```go
package algorithms

func getRow(rowIndex int) []int {
	res := make([]int, 1, rowIndex+1)
	res[0] = 1
	if rowIndex == 0 {
		return res
	}

	for i := 0; i < rowIndex; i++ {
		res = append(res, 1)
		for j := len(res) - 2; j > 0; j-- {
			res[j] += res[j-1]
		}
	}

	return res
}

```