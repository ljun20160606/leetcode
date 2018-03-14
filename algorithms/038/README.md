# Count and Say

## 描述

```txt
The count-and-say sequence is the sequence of integers with the first five terms as following:
1.     1
2.     11
3.     21
4.     1211
5.     111221



1 is read off as &#34;one 1&#34; or 11.
11 is read off as &#34;two 1s&#34; or 21.
21 is read off as &#34;one 2, then one 1&#34; or 1211.



Given an integer n, generate the nth term of the count-and-say sequence.



Note: Each term of the sequence of integers will be represented as a string.


Example 1:
Input: 1
Output: &#34;1&#34;



Example 2:
Input: 4
Output: &#34;1211&#34;


```

## 题解

```go
package algorithms

import "strconv"

//The count-and-say sequence is the sequence of integers with the first five terms as following:
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//1 is read off as "one 1" or 11.
//11 is read off as "two 1s" or 21.
//21 is read off as "one 2, then one 1" or 1211.
//Given an integer n, generate the nth term of the count-and-say sequence.
//
//Note: Each term of the sequence of integers will be represented as a string.

//Example 1:
//
//Input: 1
//Output: "1"
//Example 2:
//
//Input: 4
//Output: "1211"

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	count := 1
	var r []byte
	s := countAndSay(n - 1)
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			r = append(r, []byte(strconv.Itoa(count))...)
			r = append(r, s[i-1])
			count = 1
			continue
		}
		count++
	}
	r = append(r, []byte(strconv.Itoa(count))...)
	r = append(r, s[len(s)-1])
	return string(r)
}

```