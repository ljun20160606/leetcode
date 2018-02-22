package algorithms

//Implement wildcard pattern matching with support for '?' and '*'.
//
//'?' Matches any single character.
//'*' Matches any sequence of characters (including the empty sequence).
//
//The matching should cover the entire input string (not partial).
//
//The function prototype should be:
//bool isMatch(const char *s, const char *p)
//
//Some examples:
//isMatch("aa","a") → false
//isMatch("aa","aa") → true
//isMatch("aaa","aa") → false
//isMatch("aa", "*") → true
//isMatch("aa", "a*") → true
//isMatch("ab", "?*") → true
//isMatch("aab", "c*a*b") → false

func wildCardIsMatch(s string, p string) bool {
	lenY, lenX := len(s)+1, len(p)+1
	dp := make([][]bool, lenY)
	for y := 0; y < lenY; y++ {
		dp[y] = make([]bool, lenX)
	}
	dp[0][0] = true
	for x := 1; x < lenX; x++ {
		if p[x-1] == '*' {
			dp[0][x] = dp[0][x-1]
		}
	}
	for y := 1; y < lenY; y++ {
		for x := 1; x < lenX; x++ {
			switch p[x-1] {
			case '*':
				dp[y][x] = dp[y][x-1] || dp[y-1][x]
			case '?', s[y-1]:
				dp[y][x] = dp[y-1][x-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}

func bestWildCardIsMatch(s string, p string) bool {
	slen, plen := len(s), len(p)
	star, ss := -1, 0
	spos, ppos := 0, 0
	for spos < slen {
		// advancing both pointers when (both characters match) or ('?' is found in pattern)
		if ppos < plen && (p[ppos] == '?' || p[ppos] == s[spos]) {
			spos++
			ppos++
			continue
		}

		// * found in pattern, track index of *, only advancing pattern pointer
		if ppos < plen && p[ppos] == '*' {
			star = ppos
			ppos++
			ss = spos
			continue
		}

		// current character didn't match, last pattern pointer was '*', current pattern pointer is not '*'
		if star >= 0 {
			ppos = star + 1
			spos = ss + 1
			ss++
			continue
		}

		// current pattern pointer is not star, last pattern pointer was not '*', character do not match
		return false
	}

	for ; ppos < plen && p[ppos] == '*'; ppos++ {
	}

	return ppos == plen && spos == slen
}
