# Reverse String

## 描述

```txt
Write a function that takes a string as input and returns the string reversed.

Example 1:


Input: "hello"
Output: "olleh"



Example 2:

Input: "A man, a plan, a canal: Panama"
Output: "amanaP :lanac a ,nalp a ,nam A"




```

## 题解

```go
package algorithms

import "strings"

func reverseString(s string) string {
	builder := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		builder.WriteByte(s[i])
	}
	return builder.String()
}
```