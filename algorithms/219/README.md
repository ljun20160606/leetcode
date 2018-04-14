# Contains Duplicate II

## 描述

```txt

Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.

```

## 题解

```go
package algorithms

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if vv, has := m[v]; has {
			if i-vv <= k {
				return true
			}
		}
		m[v] = i
	}
	return false
}

```