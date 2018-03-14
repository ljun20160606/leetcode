# Letter Combinations of a Phone Number

## 描述

```txt
Given a digit string, return all possible letter combinations that the number could represent.



A mapping of digit to letters (just like on the telephone buttons) is given below.


Input:Digit string &#34;23&#34;
Output: [&#34;ad&#34;, &#34;ae&#34;, &#34;af&#34;, &#34;bd&#34;, &#34;be&#34;, &#34;bf&#34;, &#34;cd&#34;, &#34;ce&#34;, &#34;cf&#34;].



Note:
Although the above answer is in lexicographical order, your answer could be in any order you want.

```

## 题解

```go
package algorithms

//Given a digit string, return all possible letter combinations that the number could represent.
//
//A mapping of digit to letters (just like on the telephone buttons) is given below.

var table = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	var result []string
	helper(&result, digits, "", 0)
	return result
}

func helper(result *[]string, digits, prefix string, index int) {
	if index >= len(digits) {
		if len(prefix) > 0 {
			*result = append(*result, prefix)
		}
		return
	}
	digit := digits[index] - '0'
	if digit < 2 {
		helper(result, digits, prefix, index+1)
	} else {
		for _, v := range table[digit] {
			helper(result, digits, prefix+string(v), index+1)
		}
	}
}

```