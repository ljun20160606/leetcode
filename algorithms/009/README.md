# Palindrome Number

## 描述

```txt
Determine whether an integer is a palindrome. Do this without extra space.

click to show spoilers.

Some hints:

Could negative integers be palindromes? (ie, -1)

If you are thinking of converting the integer to string, note the restriction of using extra space.

You could also try reversing an integer. However, if you have solved the problem &#34;Reverse Integer&#34;, you know that the reversed integer might overflow. How would you handle such case?

There is a more generic way of solving this problem.


```

## 题解

```go
package algorithms

//Determine whether an integer is a palindrome. Do this without extra space.

// 反转数字 判断是否相等
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	t := x
	var r int
	for x > 0 {
		r = r*10 + x%10
		x /= 10
	}
	return t == r
}

```