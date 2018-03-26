# Valid Palindrome

## 描述

```txt

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.



For example,
&#34;A man, a plan, a canal: Panama&#34; is a palindrome.
&#34;race a car&#34; is not a palindrome.



Note:
Have you consider that the string might be empty? This is a good question to ask during an interview.

For the purpose of this problem, we define empty string as valid palindrome.

```

## 题解

```go
package algorithms

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && ignore(s[left]) {
			left ++
		}
		for left < right && ignore(s[right]) {
			right--
		}

		if lowerCase(s[left]) != lowerCase(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func ignore(b byte) bool {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
		return false
	}
	return true
}

func lowerCase(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 'a' - 'A'
	}
	return b
}

```