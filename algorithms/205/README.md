# Isomorphic Strings

## 描述

```txt
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

For example,
Given "egg", "add", return true.

Given "foo", "bar", return false.

Given "paper", "title", return true.

Note:
You may assume both s and t have the same length.
```

## 题解

```go
package algorithms

func isIsomorphic(s string, t string) bool {
	m1 := make([]int, 256)
	m2 := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if m1[int(s[i])] != m2[int(t[i])] {
			return false
		}

		i2 := i + 1
		m1[int(s[i])] = i2
		m2[int(t[i])] = i2
	}

	return true
}

```