# Implement strStr()

## 描述

```txt

Implement strStr().



Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.


Example 1:
Input: haystack = &#34;hello&#34;, needle = &#34;ll&#34;
Output: 2



Example 2:
Input: haystack = &#34;aaaaa&#34;, needle = &#34;bba&#34;
Output: -1


```

## 题解

```go
package algorithms

//Implement strStr().
//
//Returns the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

func strStr(haystack, needle string) int {
	ls, lss := len(haystack), len(needle)
	if ls|lss == 0 {
		return 0
	}
	end := ls - lss + 1
	for i := 0; i < end; i++ {
		var t int
		for ii := range needle {
			if haystack[i+ii] != needle[ii] {
				break
			}
			t++
		}
		if t == lss {
			return i
		}
	}
	return -1
}

```