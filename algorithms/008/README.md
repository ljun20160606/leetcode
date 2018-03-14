# String to Integer (atoi)

## 描述

```txt
Implement atoi to convert a string to an integer.

Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.

Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.

 

Requirements for atoi:

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned. If the correct value is out of the range of representable values, INT_MAX (2147483647) or INT_MIN (-2147483648) is returned.

```

## 题解

```go
package algorithms

import (
	"math"
)

//Implement atoi to convert a string to an integer.
//
//Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.
//
//Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.

// 1. 去空格
// 2. 判断正负值
// 3. 取数字
func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	// 去空格
	i := 0
	for str[i] == ' ' || str[i] == '\t' {
		i++
	}
	str = str[i:]

	// 判断正负
	flag := 1
	switch str[0] {
	case '-':
		flag = -1
		fallthrough
	case '+':
		str = str[1:]
	}

	// 只有正负号
	if len(str) == 0 {
		return 0
	}

	// 去0
	i = 0
	for str[i] == '0' {
		i++
	}

	// 取值
	result := 0
	for _, b := range str {
		n := int(b) - 48
		if n < 0 || n > 9 {
			break
		}
		result = result*10 + n
		if result > math.MaxInt32 {
			if flag == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
	}
	return result * flag
}

```