package algorithms

//Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
//
//For "(()", the longest valid parentheses substring is "()", which has length = 2.
//
//Another example is ")()())", where the longest valid parentheses substring is "()()", which has length = 4.

func longestValidParentheses(s string) int {
	var longest, substringStart int
	stk := make([]int, len(s))
	top := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' { // '('入栈
			top++
			stk[top] = i
		} else { // s[i] == ')' 两种情况
			if top == -1 { // 1. 栈为空，substringStart为下一个字符
				substringStart = i + 1
			} else { // 2. 栈不为空，两种情况
				top-- // '('出栈
				var length int
				if top == -1 { // (1) 栈为空，substring有效，length+1
					length = i - substringStart + 1
				} else { // (2) 栈不为空，取有效部分substring长度
					length = i - stk[top]
				}
				if length > longest {
					longest = length
				}
			}
		}
	}
	return longest
}

func longestValidParenthesesBest(s string) int {
	l := 0
	if len(s) >= 1 {
		i, stack := 0, make([]int, len(s)+1)
		stack[0] = -1
		for si := 0; si < len(s); si++ {
			if s[si] == '(' {
				i++ // push
				stack[i] = si
			} else {
				i-- // pop
				if i < 0 {
					i = 0 // push
					stack[i] = si
				} else if tmp := si - stack[i]; tmp > l {
					l = tmp
				}
			}
		}
	}
	return l
}
