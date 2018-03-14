# Valid Parentheses

## 描述

```txt
Given a string containing just the characters &#39;(&#39;, &#39;)&#39;, &#39;{&#39;, &#39;}&#39;, &#39;[&#39; and &#39;]&#39;, determine if the input string is valid.

The brackets must close in the correct order, &#34;()&#34; and &#34;()[]{}&#34; are all valid but &#34;(]&#34; and &#34;([)]&#34; are not.

```

## 题解

```go
package algorithms

//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.

// 看到的另一种简易栈的实现方式，很有趣
// 申请len(s)长度的[]byte skt
// 栈顶top=-1
// 入栈top+=1
// 出栈top-=1
// 取栈顶stk[top]
func isValid(s string) bool {
	stk := make([]byte, len(s))
	top := -1
	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			stk[top+1] = s[i]
			top++
		case ')':
			if top == -1 || stk[top] != '(' {
				return false
			}
			top--
		case ']':
			if top == -1 || stk[top] != '[' {
				return false
			}
			top--
		case '}':
			if top == -1 || stk[top] != '{' {
				return false
			}
			top--
		}
	}
	return top == -1
}

```