package algorithms

import "fmt"

//Implement regular expression matching with support for '.' and '*'.
//
//'.' Matches any single character.
//'*' Matches zero or more of the preceding element.
//
//The matching should cover the entire input string (not partial).
//
//The function prototype should be:
//bool isMatchV0(const char *s, const char *p)
//
//Some examples:
//isMatchV0("aa","a") ? false
//isMatchV0("aa","aa") ? true
//isMatchV0("aaa","aa") ? false
//isMatchV0("aa", "a*") ? true
//isMatchV0("aa", ".*") ? true
//isMatchV0("ab", ".*") ? true
//isMatchV0("aab", "c*a*b") ? true

// v0版本的精简版 去掉了parser那一套 清爽了很多 本质上就是 状态机+回溯
func isMatchV1(s string, p string) bool {
	fmt.Println(s, p)
	switch {
	case len(p) == 0:
		return len(s) == 0
	case len(s) == 0:
		if len(p) > 1 && p[1] == '*' {
			return isMatchV1(s, p[2:])
		}
		return false
	case len(p) > 1 && p[1] == '*':
		if isMatchV1(s, p[2:]) {
			return true
		} else if s[0] == p[0] || p[0] == '.' {
			return isMatchV1(s[1:], p)
		}
		return false
	default:
		return (s[0] == p[0] || p[0] == '.') && isMatchV1(s[1:], p[1:])
	}
}

// dp
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}
	for i, mv := range s {
		for j, pv := range p {
			switch pv {
			case '*':
				if p[j-1] == s[i] || p[j-1] == '.' {
					dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j] || dp[i+1][j-1]
				} else {
					dp[i+1][j+1] = dp[i+1][j-1]
				}
			case '.', mv:
				dp[i+1][j+1] = dp[i][j]
			}
		}
	}
	return dp[m][n]
}
