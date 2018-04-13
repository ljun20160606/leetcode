# Contains Duplicate

## 描述

```txt

Given an array of integers, find if the array contains any duplicates. Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

```

## 题解

```go
package algorithms

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		} else {
			m[n] = struct{}{}
		}
	}
	return false
}

```