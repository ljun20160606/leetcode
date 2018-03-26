# Length of Last Word

## 描述

```txt
Given a string s consists of upper/lower-case alphabets and empty space characters &#39; &#39;, return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:
Input: &#34;Hello World&#34;
Output: 5


```

## 题解

```go
package algorithms

func lengthOfLastWord(s string) int {
	var r int
	for e := range s {
		if s[e] != ' ' {
			if e > 0 && s[e-1] == ' '{
				r = 0
			}
			r++
		}
	}
	return r
}

```