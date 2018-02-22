package algorithms

//Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
//
//Example:
//
//Input: "babad"
//
//Output: "bab"
//
//Note: "aba" is also a valid answer.
//Example:
//
//Input: "cbbd"
//
//Output: "bb"

// 增加一个字符后有三种可能
// 1. 回环长度+1 如 a + a
// 2. 回环长度+2 如 ab + a
// 3. 回环长度不变 如 a + b
// 起始位置start=0 回环长度maxLen=1
// 截取末尾 maxLen+2长度字符 如果对称 则回环长度+2 变更起始位置
// 截取末尾 maxLen+1长度字符 如果对称 则回环长度+1 变更起始位置
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start := 0
	maxLen := 1
	for index := range s {
		l := index - maxLen
		end := index + 1
		if l >= 1 && sym(s[l-1:end]) {
			start = l - 1
			maxLen += 2
		} else if l >= 0 && sym(s[l:end]) {
			start = l
			maxLen += 1
		}
	}
	return s[start : start+maxLen]
}

// 当odd长度时对比中轴前后字符
// 当even长度时对比中轴前后字符
// 时间复杂度O(n/2)
// 当然还能选择反转字符对比字符串是否相等，但是相应的时间复杂度会变成 O(n)
func sym(s string) bool {
	end := len(s) - 1
	if len(s)&1 == 1 {
		mid := len(s) / 2
		for i := 0; i < mid; i++ {
			if s[i] != s[end-i] {
				return false
			}
		}
	} else {
		mid := len(s)/2 - 1
		for i := 0; i <= mid; i++ {
			if s[i] != s[end-i] {
				return false
			}
		}
	}
	return true
}
