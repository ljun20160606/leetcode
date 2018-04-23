# Group Anagrams

## 描述

```txt
Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]

Note:


	All inputs will be in lowercase.
	The order of your output does not matter.


```

## 题解

```go
package algorithms

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		bytes := []byte(v)
		sort.SliceStable(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		s := string(bytes)
		m[s] = append(m[s], v)
	}
	var ss [][]string
	for e := range m {
		ss = append(ss, m[e])
	}
	return ss
}

```