# Integer to Roman

## 描述

```txt
Given an integer, convert it to a roman numeral.


Input is guaranteed to be within the range from 1 to 3999.
```

## 题解

```go
package algorithms

//Given an integer, convert it to a roman numeral.
//
//Input is guaranteed to be within the range from 1 to 3999.

// 还有一种骚操作，不迭代，直接/1000 /100 /10
func intToRoman(num int) string {
	roman := [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}

	var result string
	var flag int
	for num > 0 {
		t := num % 10
		result = roman[flag][t] + result
		flag++
		num /= 10
	}
	return result
}

```