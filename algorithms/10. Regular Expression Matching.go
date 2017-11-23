package algorithms

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
	lenY, lenX := len(s)+1, len(p)+1
	dp := make([][]bool, lenY)
	for y := range dp {
		dp[y] = make([]bool, lenX)
	}
	dp[0][0] = true
	for x := 1; x < lenX; x++ {
		if p[x-1] == '*' {
			dp[0][x] = dp[0][x-2]
		}
	}
	for y := 1; y < lenY; y++ {
		for x := 1; x < lenX; x++ {
			switch p[x-1] {
			case '*':
				if p[x-2] == s[y-1] || p[x-2] == '.' {
					dp[y][x] = dp[y-1][x] || dp[y][x-1] || dp[y][x-2]
				} else {
					dp[y][x] = dp[y][x-2]
				}
			case '.', s[y-1]:
				dp[y][x] = dp[y-1][x-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}
