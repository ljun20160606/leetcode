# Longest Substring Without Repeating Characters

## 描述

```txt
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given &#34;abcabcbb&#34;, the answer is &#34;abc&#34;, which the length is 3.

Given &#34;bbbbb&#34;, the answer is &#34;b&#34;, with the length of 1.

Given &#34;pwwkew&#34;, the answer is &#34;wke&#34;, with the length of 3. Note that the answer must be a substring, &#34;pwke&#34; is a subsequence and not a substring.
```

## 题解

```go
package algorithms

//Given a string, find the length of the longest substring without repeating characters.
//
//Examples:
//
//Given "abcabcbb", the answer is "abc", which the length is 3.
//
//Given "bbbbb", the answer is "b", with the length of 1.
//
//Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

// 根据ascii码特性, 建立字母表dict
// start: 未重复substr起始位置, lenMax: 最大长度
// 将字母位置放入字母表
// 字母表中字母位置大于start代表字母重复出现, 将start右移至重复出现的位置
// 根据start和当前index计算长度
func lengthOfLongestSubstring(s string) int {
	start := -1
	lenMax := 0
	dict := [256]int{}
	for i := range dict {
		dict[i] = -1
	}
	for i, r := range s {
		if v := dict[r]; v > start {
			start = v
		}
		length := i - start
		if length > lenMax {
			lenMax = length
		}
		dict[r] = i
	}
	return lenMax
}

```